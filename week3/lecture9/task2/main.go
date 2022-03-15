package main

import (
	"errors"
	"fmt"
)

type Action func() error

func noError() error {
	return nil
}

func errorFunc() error {
	return errors.New("first error")
}

func panicFunc() error {
	panic("panic happened")
}

func SafeExec(a Action) Action {

	b := func() error {

		err := a()
		return fmt.Errorf("safe exec: %w", err)
	}

	return b
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	res1 := SafeExec(noError)
	res2 := SafeExec(errorFunc)
	res3 := SafeExec(panicFunc)
	fmt.Println(res1())
	fmt.Println(res2())
	fmt.Println(res3())

}
