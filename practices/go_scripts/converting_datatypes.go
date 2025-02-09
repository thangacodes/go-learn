package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("First converting integer to float..\n")
	var i int = 40
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("Sleep for '1' seconds..\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("Secondly converting float to integer..\n")
	var fl float64 = 100.2343
	var in int = int(fl)
	fmt.Printf("%v\n", in)
}
