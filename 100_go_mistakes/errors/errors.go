package main

import (
	"errors"
	"fmt"
	"reflect"
)

type ErrorType string

const (
	ERROR_GENERIC        ErrorType = "generic_error"
	ERROR_INVALID_DATA   ErrorType = "invalid_data"
	ERROR_MISSING_DATA   ErrorType = "missing_data"
	ERROR_DUPLICATE_DATA ErrorType = "duplicate_data"
	ERROR_DATA_NOT_FOUND ErrorType = "data_not_found"
	ERROR_DATABASE       ErrorType = "database_error"
	ERROR_NOT_ALLOWED    ErrorType = "not_allowed"
	ERROR_UNAUTHORISED   ErrorType = "unauthorised"
	ERROR_FORBIDDEN      ErrorType = "forbidden"
	ERROR_IS_DELETED     ErrorType = "is_deleted"
)

func SafeSprintf(msg string, a ...any) string {
	args := make([]any, 0)

	for _, arg := range a {
		argValue := reflect.ValueOf(arg)

		if argValue.Kind() == reflect.Ptr {
			if argValue.IsNil() {
				args = append(args, "nil")
			} else {
				args = append(args, argValue.Elem())
			}
		} else {
			args = append(args, argValue)
		}
	}

	return fmt.Sprintf(msg, args...)
}

type DCError interface {
	ErrorType() ErrorType
	Error() string
}

type dcError struct {
	errorType ErrorType
	msg       string
}

func (err *dcError) ErrorType() ErrorType {
	return err.errorType
}

func (err *dcError) Error() string {
	return err.msg
}

func newDCError(errorType ErrorType, msg string, a ...any) error {
	errMsg := SafeSprintf(msg, a...)
	return &dcError{errorType: errorType, msg: errMsg}
}

func NewMissingDataError(msg string, a ...any) error {
	return newDCError(ERROR_MISSING_DATA, msg, a...)
}

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

	// Joined errors never get wrapped
	joinedError := errors.Join(myMainError, fmt.Errorf("this is another error"))
	fmt.Println(joinedError)
	fmt.Println(errors.Is(myMainError, joinedError))

	fmt.Println()
	fmt.Println("Testing Custom error unwrap")
	missingDataError := NewMissingDataError("this is a missing data error")
	fmt.Println("error not wrapped, testing errors.As function")
	var dcErr DCError
	fmt.Println("is error found with errors.As?", errors.As(missingDataError, &dcErr))

	fmt.Println()
	fmt.Println("Testing Custom error wrapped")
	missindDataErrW := NewMissingDataError("this is a missing data error")
	missDataWrapped := fmt.Errorf("this is a wrapped error: %w", missindDataErrW)
	fmt.Println("error wrapped, testing errors.As function")
	var dcErrW DCError
	fmt.Println("is error found with errors.As?", errors.As(missDataWrapped, &dcErrW))
	fmt.Println("is error found with errors.Is?", errors.Is(missDataWrapped, dcErrW))
}

func doSomethingWithErr() error {
	return myMainError
}

func wrapTheError() error {
	err := doSomethingWithErr()
	if err != nil {
		return fmt.Errorf("something failed: %w", myMainError) // wrapped
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
