package main

import (
	"fmt"
	"sync"
)

func main() {

	/*
		Mutex is a way to synchronize the way we call a function, for example, whenever a function is in "use" by a
		go routine we lock it we finish our process to let other process to go into that function
	*/

	var balance int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 102; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deposit(&balance, 1, &mu)
		}()
	}

	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(&balance, 1, &mu)
		}()
	}

	wg.Wait()
	fmt.Println(balance)

}

/* Created deposit and withdrawal methods to only add or subtract an exact amount
notice how we are using the lock and unlock methods of mutex
NOTE: We pass balance's pointer to avoid a return
*/

func deposit(balance *int, amount int, mu *sync.Mutex) {
	mu.Lock()
	*balance += amount
	defer mu.Unlock() // Use defer to make sure we unlock whenever we finish all the steps in the function
}

func withdrawal(balance *int, amount int, mu *sync.Mutex) {
	mu.Lock()
	*balance -= amount
	defer mu.Unlock()
}
