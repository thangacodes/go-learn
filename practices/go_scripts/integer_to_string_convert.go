package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Converting Integer to String using 'Itoa()' function..\n")
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)
	fmt.Printf("Sleep for '1' seconds..\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("Converting String to Integer using 'Atoi()' function..\n")
	var st string = "2000"
	in, err := strconv.Atoi(st)
	fmt.Printf("%v %T\n", in, in)
	fmt.Printf("%v %T\n", err, err)
}
