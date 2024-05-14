package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testingTheErrors()
	fmt.Println()
	fmt.Println(err)
}

func testingTheErrors() (err error) {
	res, err := thisMayThrowErrorOne(10)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := thisIsDeferError(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
			} else {
				err = closeErr
			}
			fmt.Println(err)
		}
	}()

	res, err = thisMayThrowErrorOne(15)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func thisMayThrowErrorOne(a int) (int, error) {
	if b := a % 2; b != 0 {
		return 0, fmt.Errorf("number not divisible by even")
	} else {
		return b, nil
	}
}

func thisIsDeferError() error {
	return fmt.Errorf("this is the deferred error")
}
