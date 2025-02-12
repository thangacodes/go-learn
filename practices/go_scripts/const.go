package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Const in Go program")
	var num = 10
	var country = "india"
	fmt.Printf("Given variable 'num' datatype of: %s\n", reflect.TypeOf(num))
	fmt.Printf("Given variable 'country' datatype of: %s\n", reflect.TypeOf(country))
	time.Sleep(1 * time.Second)
	const name = "nancy"
	fmt.Printf("The name of given variable 'name' belongs to datatype of: %v: %T \n", name, name)
}

// %v prints the value.
// %T prints the type of the variable.
// You need to pass the value twice, once for each format specifier:
// one for the value and one for the type.
