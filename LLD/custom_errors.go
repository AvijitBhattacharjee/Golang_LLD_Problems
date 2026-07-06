package main

import (
	"errors"
	"fmt"
)


type customUserError struct {
	id int
}


func (c *customUserError) Error() string {
	return fmt.Sprintf("user %d crashed ", c.id)
}

func svc(id int) error {

	if id > 10 {
		return &customUserError{id: id}
	}
	return nil
}

func main() {

	// storing error 
	err := errors.New("new error")
	fmt.Println(err)

	// another way to store error 
	err = fmt.Errorf("user %s not found", "avijit")
	fmt.Println(err)

	// error wrapping
	wrapped_error := service()
	fmt.Println("wrapper error due to :", wrapped_error)

	// error unwrapping
	unwrapped_error := errors.Unwrap(wrapped_error)
	fmt.Println("one level down error = ", unwrapped_error)

	// errors.As is used when you want to extract a specific error type from a wrapped error chain.
	// 
	// %v only formats the error into text.
	// %w wraps the error and preserves the original error for errors.Is, errors.As, and errors.Unwrap.
	if errors.Is(wrapped_error, errors.New("service error due to database error:")) {
		fmt.Println("Hi")
	}

	// custom error
	fmt.Println(svc(200))
	
}

func service() error {
	err := database()
	return fmt.Errorf("service error due to %w:", err)
}

func database() error {
	return errors.New("database error")
}