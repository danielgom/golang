package main

import (
	"errors"
	"fmt"
)

var myMainError = errors.New("this is the main error")

type this struct {
	first  string
	second string
}

func main() {
	mainError := doSomethingWithErr()
	fmt.Println(mainError)

	wrappedError := wrapTheError()
	fmt.Println(errors.Unwrap(wrappedError))

	fmt.Println(errors.Is(errors.Unwrap(wrappedError), mainError)) // true
	fmt.Println(errors.Is(wrappedError, mainError))                // true

	doubleWrappedError := doubleWrapError()

	// Even if we wrap multiple times, the error is still the main error
	fmt.Println(errors.Is(errors.Unwrap(doubleWrappedError), mainError)) // true
	fmt.Println(doubleWrappedError)
	fmt.Println(mainError)
	fmt.Println(errors.Is(doubleWrappedError, mainError)) // true

	unwrappedError := noWrapTheError()
	fmt.Println(errors.Unwrap(unwrappedError)) // not wrapped, nothing to be shown

	fmt.Println(errors.Is(errors.Unwrap(unwrappedError), mainError)) // false
	fmt.Println(errors.Is(unwrappedError, mainError))                // false
}

func doSomethingWithErr() error {
	return myMainError
}

func wrapTheError() error {
	err := doSomethingWithErr()
	if err != nil {
		return fmt.Errorf("something failed: %w", err) // wrapped
	}

	return err
}

func doubleWrapError() error {
	err := wrapTheError()
	if err != nil {
		return fmt.Errorf("something failed again: %w", err)
	}

	return err
}

func noWrapTheError() error {
	err := doSomethingWithErr()
	if err != nil {
		return fmt.Errorf("something failed: %v", err) // not wrapped
	}

	return err
}
