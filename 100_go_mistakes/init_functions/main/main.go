package main

import (
	"fmt"
	"github.com/dgomez/learngo/first/explain/100_go_mistakes/init_functions/redis"
	"log"
	"sync"
)

// We should not rely on init functions order
func main() {
	// Last execution
	fmt.Println("here we are")
	err := redis.Store("foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	// Avoiding init functions is neither a good nor a bad practice, however
	// it may lead to mistakes, using sync.Once function and derivatives is
	// greatly advised

	onceFunc := sync.OnceFunc(shouldInitializeOnce)
	onceFunc()
	onceFunc() // Just runs once
}

func init() {
	fmt.Println("init 1") // First to be executed
}

func init() {
	fmt.Println("init 2") // Second to be executed
}

func shouldInitializeOnce() {
	fmt.Println("Do something")
}
