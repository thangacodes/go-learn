package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string
	var b int
	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("Enter variable 'a' is datatype of: ", reflect.TypeOf(a))
	fmt.Println("Enter variable 'b' is datatype of: ", reflect.TypeOf(b))

}
