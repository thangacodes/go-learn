package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	fmt.Printf("%d\n", i)
	fmt.Printf("Sleep for '2' seconds..\n")
	time.Sleep(2 * time.Second)
	var f1 float64
	fmt.Printf("%.2f\n", f1)
	fmt.Printf("Sleep for '2' seconds..\n")
	time.Sleep(2 * time.Second)
	var cj bool
	fmt.Printf("%t\n", cj)
	fmt.Printf("Sleep for '1' seconds..\n")
	time.Sleep(1 * time.Second)
	var foo string
	fmt.Printf("%s\n", foo)
}
