package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "India is my country"
	fmt.Println(a)
	fmt.Println("Type of 'a' is:", reflect.TypeOf(a))
	b := 76534543
	fmt.Println(b)
	fmt.Println("Type of 'b' is:", reflect.TypeOf(a))
	c := true
	fmt.Println(c)
	fmt.Println("Type of 'c' is:", reflect.TypeOf(c))
}

// reflect.TypeOf(variable) --> returns the type of the variable, which is then printed out.
