package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("..........")
	fmt.Printf("Example program for: equal(==),not equal(!=),less than(<),less than or equal to(<=), greater than(>), greater than(>=)\n")
	fmt.Println("")
	fmt.Println("..........")
	var x int = 100
	var y int = 100
	fmt.Println("var1('100') is equal to var2('100'):", x == y)
	var city string = "chennai"
	var city1 string = "madras"
	fmt.Println("var1('chennai') is not equal to var2('madras') and the answer is:", city != city1)
	time.Sleep(1 * time.Second)
	var a, b int = 5, 6
	fmt.Println("var1('5') is less than var2('6'):", a < b)
	time.Sleep(1 * time.Second)
	var c, d int = 100, 100
	fmt.Println("var1('100') is less than or equal to var2('100'):", c <= d)
	time.Sleep(1 * time.Second)
	var e, f int = 200, 20
	fmt.Println("var1('200') is greater than var2('20'):", e > f)
	time.Sleep(1 * time.Second)
	var g, i int = 200, 200
	fmt.Println("var1('200') is greater than or equal to var2('200'):", g >= i)
}
