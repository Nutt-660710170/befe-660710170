package main

import (
	"fmt"
	"errors"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

func main() {

	result, error:= divide(10,0)
	if error != nil {
		fmt.Println("Error",error)
	}
	fmt.Println("Result = ", result)
}