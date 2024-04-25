package main

import (
	"context"
	"errors"
	"golang.org/x/sync/semaphore"
	"log"
	"math/rand/v2"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	perm := rand.Perm(2 << 19)
	permString := make([]string, 0, len(perm))
	for _, p := range perm {
		permString = append(permString, strconv.Itoa(p))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	producerChan, err := producer(ctx, permString)
	if err != nil {
		log.Fatal(err)
	}

	step1results, step1errors := step(ctx, producerChan, transformA)
	step2results, step2errors := step(ctx, step1results, transformB)
	allErrors := merge(ctx, step1errors, step2errors)

	sink(ctx, step2results, allErrors)
}

func producer(ctx context.Context, strings []string) (<-chan string, error) {
	outChan := make(chan string)

	go func() {
		defer close(outChan)
		for _, s := range strings {
			select {
			case <-ctx.Done():
				return
			case outChan <- s:
				log.Println("sent to the channel", s)
			}
		}
	}()

	return outChan, nil
}

func sink(ctx context.Context, values <-chan string, errors <-chan error) {
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err().Error())
			return
		case err := <-errors:
			if err != nil {
				log.Println("error", err.Error())
				//cancel()
			}
		case val, ok := <-values:
			if !ok {
				log.Println("Done")
				return
			}
			log.Println("Sink:", val)
		}
	}
}

func step[In any, Out any](ctx context.Context, inChan <-chan In, fn func(In) (Out, error)) (chan Out, chan error) {
	outputChannel := make(chan Out)
	errorChannel := make(chan error)

	limit := int64(runtime.NumCPU())
	sem := semaphore.NewWeighted(limit)

	go func() {
		defer close(outputChannel)
		defer close(errorChannel)
		for {
			select {
			case <-ctx.Done():
				break
			case val, ok := <-inChan:
				if ok {
					if err := sem.Acquire(ctx, 1); err != nil {
						log.Println("Failed to acquire semaphore:", err)
						break
					}
					go func(s In) {
						defer sem.Release(1)
						result, err := fn(s)
						if err != nil {
							errorChannel <- err
						} else {
							outputChannel <- result
						}
					}(val)
				} else {
					if err := sem.Acquire(ctx, limit); err != nil {
						log.Println("Failed to acquire semaphore:", err)
					}
					return
				}
			}
		}
	}()

	return outputChannel, errorChannel
}

func merge[T any](ctx context.Context, cs ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	output := func(c <-chan T) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out // we can always return a closed channel to read.
}

func transformA(s string) (string, error) {
	log.Println("transform A input:", s)
	if s == "FOO" {
		return "", errors.New("something happened :C")
	}
	return strings.ToLower(s), nil
}

func transformB(s string) (string, error) {
	log.Println("transform B input: ", s)

	return strings.ToTitle(s), nil
}
