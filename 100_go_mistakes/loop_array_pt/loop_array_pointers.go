package main

import "fmt"

type Customer struct {
	ID      int64
	Name    string
	Balance float64
}

type Store struct {
	m map[int64]*Customer
}

func (s *Store) storeCustomers(customers []*Customer) {
	if s.m == nil {
		s.m = make(map[int64]*Customer)
	}
	// After go 1.22
	for _, customer := range customers {
		// Careful here, sharing pointer of array and map
		fmt.Printf("Customer pointer: %p\n", customer)
		s.m[customer.ID] = customer
	}

	// Before go 1.22
	/*
		for _, customer := range customers {
			current := customer
			s.m[customer.ID] = &current
		}

	*/
}

func (s *Store) checkCustomers() {
	for _, c := range s.m {
		fmt.Printf("%+v with pointer %p\n ", c, c)
	}
}

func main() {
	myStore := &Store{}

	customerList := []*Customer{
		{
			ID:      1,
			Name:    "John Doe",
			Balance: 42.0,
		},
		{
			ID:      2,
			Name:    "Jane Doe",
			Balance: 51.8,
		},
		{
			ID:      3,
			Name:    "Dan Johnson",
			Balance: 12.34,
		},
	}
	myStore.storeCustomers(customerList)

	fmt.Printf("%p\n", &customerList[1])

	// Shared pointer, anything changed in the array will affect the map
	customerList[1].Balance = 100.25

	// Affected map after change
	myStore.checkCustomers()

}
