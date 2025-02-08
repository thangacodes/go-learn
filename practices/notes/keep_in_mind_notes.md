```bash
Datatypes:
int, string, boolean
int has two types - uint, int
unit: (unsigned integer)
it contains only positive numbers, including '0' also.
int: (signed integer)
it contains both positive and negative numbers. 

Printf -format specifiers:
verb Descriptions
%v   formats the value in a default format.
%d   formats decimal integers
%T   type of the value
%c   character
%q   quoted characters/string
%s   plain string
%t   true or false
%f   floating numbers
%.2f floating numbers upto 2 decimal places

## Variable Declaration Methods:
   Method I and Method II are shorthand ways to declare variables:

* Method I:
  var s, t = "foo", "bar"

* Method II:
  var (
    name = "John"
    age  = 35
    height = 5.7
    weight = 67.8
    married = false 
)

* Short Variable Declaration:
  s := "India"   //:= denotes short variable declaration

# variables_dec.go
package main

import (
	"fmt"
)

func main() {
	name := "Arjun"
	name = "Siju"
	fmt.Println(name)
}

# When you run the program with:
  go run variables_dec.go
# The output will be:
  Siju
*  Note: You can override string variable values.

Another method is,
# variables_dec.go
package main

import (
	"fmt"
)

func main() {
	name := "Arjun"
	name = 590
	fmt.Println(name)
}

# When you run the program with:
  go run variables_dec.go

# You will encounter an error message:
Error: cannot use 590 (type untyped int) as type string in assignment.

** Note: You can override string variable values, but you cannot assign an integer to a string variable.

# Local vs Global Variables:

Local Variables:
* Declared inside a function or a block
* not accessible outside the function or the block
* can also be declared inside looping and conditional statements

Example for local variables:
# main.go 
package main
import (
	"fmt"
)

func main() {
	city := "London"
	fmt.Println(city)
}

The output will be:
London 

Global Variables:
* Declared outside of a function or  a block
* Available throughout the lifetime of a program
* Declared at the top of the program outside all functions or blocks
* Can be accessed from any part of the program

Example for Global Variables:
# main.go
package main

import (
	"fmt"
)

var name = "Vedhu"

func main() {
	var age = 35
	fmt.Println(name)
	fmt.Println(age)
}
The output will be:
Vedhu 


  

