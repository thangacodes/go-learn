package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Going to see first method (short hand) assigning variables with their values: ")
	var s, t = "john", "morris"
	fmt.Println("Variables 's & t' values are: ", s, t)
	fmt.Println("Variable 's' datatype is: ", reflect.TypeOf(s))
	fmt.Println("Variable 't' datatype is: ", reflect.TypeOf(t))
	fmt.Println("Sleep for '2' seconds\n ")
	time.Sleep(2 * time.Second)
	fmt.Println("Going to see second method (short hand) assigning variables with their values: ")
	var (
		name = "John"
		age  = 35
	)
	fmt.Println("Variables 'name & age' values are: ", name, age)
	fmt.Println("Variable 'name' datatype is: ", reflect.TypeOf(name))
	fmt.Println("Variable 'age' datatype is: ", reflect.TypeOf(age))
	fmt.Println("Sleep for '2' seconds\n ")
	time.Sleep(2 * time.Second)
	fmt.Println("Going to see third method (short variable declaration) assigning variables with their values: ")
	a := 120.2345
	b := true
	c := 5460
	d := "Go language"
	fmt.Println("Variable 'a' value is: ", a)
	fmt.Println("Variable 'a' datatype is: ", reflect.TypeOf(a))
	fmt.Println("Variable 'b' value is: ", b)
	fmt.Println("Variable 'b' datatype is: ", reflect.TypeOf(b))
	fmt.Println("Variable 'c' value is: ", c)
	fmt.Println("Variable 'c' datatype is: ", reflect.TypeOf(c))
	fmt.Println("Variable 'd' value is: ", d)
	fmt.Println("Variable 'd' datatype is: ", reflect.TypeOf(d))

}
