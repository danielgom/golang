package main

import (
	"fmt"
	"sync"
	"time"
)

// job represents a task to be executed by the worker
type job struct {
	ID int
}

type workerPool struct {
	numberOfWorkers int
	jobQueue        chan job
	results         chan int
	wg              sync.WaitGroup
}

func main() {
	numberOfWorkers := 10
	numberOfJobs := 10

	wPool := newWorkerPool(numberOfWorkers, numberOfJobs)
	for i := range numberOfJobs {
		wPool.addJob(job{ID: i})
	}

	close(wPool.jobQueue)

	wPool.start()
	wPool.wait()
	wPool.collectResults()
}

func newWorkerPool(numberOfWorkers, jobQueueSize int) *workerPool {
	return &workerPool{
		numberOfWorkers: numberOfWorkers,
		jobQueue:        make(chan job, jobQueueSize),
		results:         make(chan int, jobQueueSize),
	}
}

func (w *workerPool) worker(id int) {
	defer w.wg.Done()
	for j := range w.jobQueue {
		fmt.Println("Worker", id, "started job", j.ID)
		time.Sleep(time.Second) // simulate the job
		fmt.Println("Worker", id, "finished job", j.ID)
		w.results <- j.ID
	}
}

func (w *workerPool) start() {
	for i := range w.numberOfWorkers {
		w.wg.Add(1)
		go w.worker(i)
	}
}

func (w *workerPool) wait() {
	w.wg.Wait()
	close(w.results)
}

func (w *workerPool) addJob(job job) {
	fmt.Println("Adding job", job.ID)
	w.jobQueue <- job
}

func (w *workerPool) collectResults() {
	for result := range w.results {
		fmt.Println("Result:", result)
	}
}
