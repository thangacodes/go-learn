package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "India is my country"
	fmt.Println("Given input of 'a' value is:", a)
	fmt.Println("Type of 'a' is:", reflect.TypeOf(a))
	b := 76534543
	fmt.Println("Given input of 'b' value is:", b)
	fmt.Println("Type of 'b' is:", reflect.TypeOf(a))
	c := true
	fmt.Println("Given input of 'c' value is:", c)
	fmt.Println("Type of 'c' is:", reflect.TypeOf(c))
}

// reflect.TypeOf(variable) --> returns the type of the variable, which is then printed out.
