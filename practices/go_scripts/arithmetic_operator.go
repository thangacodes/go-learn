package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("These are the common Arithmetic Operators :\n")
	fmt.Println("addition operator(+)\nsubtraction operator (-)\nmultiplication operator(*)\ndivision operator(/)\nmodulus operator(%)\nincrement operator(++)\ndecrement operator(--)")
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("'addition(+) operator' example")
	var a, b string = "foo", "bar"
	fmt.Println(a + b)
	time.Sleep(1 * time.Second)
	fmt.Println("")
	fmt.Println("'subtraction(-) operator' example")
	var c, d float64 = 200.20, 150.24
	fmt.Printf("%.4f", c-d)
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("")
	fmt.Println("'Multiplication(*) operator' example")
	var e, f int = 30, 2
	fmt.Println(e * f)
	time.Sleep(1 * time.Second)
	fmt.Println("")
	fmt.Println("'Division(/) operator' example")
	var g, h int = 100, 2
	fmt.Println(g * h)
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("'Modulus(%) operator' example")
	var i, j int = 24, 3
	fmt.Println(i % j)
	fmt.Println("")
	time.Sleep(1 * time.Second)
	fmt.Println("'Increment(++) operator' example")
	var k int = 2
	k++
	fmt.Println(k)
	fmt.Println()
	time.Sleep(1 * time.Second)
	fmt.Println("'Decrement(--) operator' example")
	var l int = 2
	l--
	fmt.Println(l)
}
