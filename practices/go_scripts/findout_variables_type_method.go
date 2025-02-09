package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Finding variable data types using the type-1 method..")
	var age int = 35
	var name string = "John Smith"
	var ismarried bool = true
	var height float32 = 5.6
	fmt.Printf("variable age = %v is datatype of %T \n", age, age)
	fmt.Printf("variable name = %s is datatype of %T \n", name, name)
	fmt.Printf("variable ismarried = %t is datatype of %T \n", ismarried, ismarried)
	fmt.Printf("variable height = %f is datatype of %T \n", height, height)
	fmt.Println("Sleep for '1' second..")
	time.Sleep(1 * time.Second)
	fmt.Println("Finding variable data types using the type-2 method..")
	fmt.Printf("variable age = %v is datatype of %v \n", age, reflect.TypeOf(age))
	fmt.Printf("variable name = %v is datatype of %v \n", name, reflect.TypeOf(name))
	fmt.Printf("variable ismarried = %v is datatype of %v \n", ismarried, reflect.TypeOf(ismarried))
	fmt.Printf("variable height = %v is datatype of %v \n", height, reflect.TypeOf(height))
}

# Note:
   # The %v is a format specifier used in fmt.Printf in Go. It is used to print the value of a variable in a default format.
   # The %s is another format specifier in fmt.Printf, but this one is used specifically to print a string.
   # The 'reflect package' allows you to inspect types at runtime.
   # reflect.TypeOf(variable) returns the reflect.Type of the variable
