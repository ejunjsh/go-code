package main

import "fmt"

type doError struct {}

func (e *doError) Error() string{
    return ""
}

func do() error {   // error(*doError, nil)
	var err *doError
	return err  // nil of type *doError
}

func main() {
	err := do()
	fmt.Println(err == nil)
}

