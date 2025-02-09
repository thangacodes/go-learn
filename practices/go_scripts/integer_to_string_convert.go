package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Converting Integer to String using 'Itoa()' function..\n")
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)
}
