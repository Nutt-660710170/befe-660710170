package main

import (
	"fmt"
)

func main() {
	//var name string = "nutt"
	var age int = 20

	email := "nutt5346789@gmail.com"
	gpa := 3.01

	firstname, lastname := "nutt", "singsatit"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n",firstname,lastname,age,email,gpa)
}