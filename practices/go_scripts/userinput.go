package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("User input prompt for single and then stored in a variable called 'name' ")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Entered variable 'name' is datatype of: ", reflect.TypeOf(name))
	fmt.Println("Hey", name)
	fmt.Println("Sleep for '1' seconds..")
	time.Sleep(1 * time.Second)
	fmt.Println("User input prompt for multiple and then stored in variables called 'date & married' ")
	var date int
	var married bool
	fmt.Print("Enter your date of birth & Are you married or single: ")
	fmt.Scanf("%d %t", &date, &married)
	fmt.Println("User entered variables like 'date & married' values are :", date, married)
	fmt.Println("Entered variable 'date' is datatype of: ", reflect.TypeOf(date))
	fmt.Println("Entered variable 'married' is datatype of: ", reflect.TypeOf(married))

}
