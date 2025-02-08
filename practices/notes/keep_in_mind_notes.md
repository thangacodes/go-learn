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

# Note: You can override string variable values, but you cannot assign an integer to a string variable.

# variables_dec.go
package main
import ("fmt")
func main() {
   name := "Arjun"
   name  = "Siju"
fmt.Println(name)

# When you run the program with:
  go run variables_dec.go
# The output will be:
  Siju
Another method is,
# variables_dec.go
package main
import ("fmt")
func main() {
   name := "Arjun"
   name  = 590
fmt.Println(name)

# When you run the program with:
  go run variables_dec.go
# You will encounter an error message:
Error: cannot use 590 (type untyped int) as type string in assignment.




  

