package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Converting integer to string datatype..\n")
	var i int = 40
	var s string = strconv.Itoa(i)
	fmt.Printf("%q\n", s)
	fmt.Println("Converted value type is of: ", reflect.TypeOf(s))
	time.Sleep(1 * time.Second)
	fmt.Printf("Converting string to integer datatype..\n")
	var st string = "2000"
	in, err := strconv.Atoi(st)
	fmt.Printf("%v %T\n", in, in)
	fmt.Printf("%v %T\n", err, err)
	fmt.Println("Converted value type is of: ", reflect.TypeOf(in))
	time.Sleep(1 * time.Second)
	fmt.Printf("Converting integer to float datatype..\n")
	var j int = 40
	var f float64 = float64(j)
	fmt.Printf("%.2f\n", f)
	fmt.Println("Converted value type is of: ", reflect.TypeOf(f))

}
