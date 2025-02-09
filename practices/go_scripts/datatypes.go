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
	d := 88.90
	fmt.Println("Given input of 'd' value is:", d)
	fmt.Println("Type of 'c' is:", reflect.TypeOf(d))
	f := "How is your health"
	g := "John"
	fmt.Println("Combining variables like 'f' and 'g':", f, g, "?")
	fmt.Println("Combining variables like 'f' and 'g':", f, " ", g, "?")
	h := "Kodekloud"
	i := "GO language program instructor"
	fmt.Print("One of the best online learning platform is: ", h, "\n")
	fmt.Print("Priyanka is one of the best ", "", i, "\n")
	j := "Kodekloud"
	k := "GO language program instructor"
	fmt.Println("One of the best online learning platform is:", j)
	fmt.Println("Priyanka is one of the best", k)
}

// reflect.TypeOf(variable) returns the type of the variable.
// "\n" is used for a manual newline when using fmt.Print.
// fmt.Println automatically includes a newline, so no need to add "\n".
