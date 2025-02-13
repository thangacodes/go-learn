package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("Going to see 'Logical AND Operator (&&)'...")
	var x int = 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 300) && (x < 0))
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("Going to see 'Logical OR Operator (||)'...")
	var a int = 10
	fmt.Println((a < 0) || (a < 200))
	fmt.Println((a < 0) || (a > 200))
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("Going to see 'Logical NOT Operator (!)'...")
	var c, d int = 10, 20
	fmt.Println(!(c > d))
	fmt.Println(!(true))
	fmt.Println(!(false))

}
