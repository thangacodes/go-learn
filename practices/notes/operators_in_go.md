Operators:

```bash
Operators are the foundation of any programming language.
Operators allow us to perform different kinds of operations on operands.
In the Go language, operators Can be categorized based on their different functionality:

  * Arithmetic Operators
  * Relational Operators
  * Logical Operators
  * Bitwise Operators
  * Assignment Operators
  * Misc Operators


Note: We can define Operators in symbols.

Example, when we say 
a + b
// a, b -> are operands
// + -> Operators

In Go lang, we have '5' types of Operators:
1) Comparison Operators
   Example: == != < <= > >=
    equal ==
    not equal !=
    less than <
    less than or equal to <=
    greater than >
    greater than or equal to >=

2) Arithmetic Operators
   Example:  + - * / % ++ --
3) Assignment Operators
   Example: = += -= *= /= %=
4) Bitwise Operators
   Example: & | <<  >> ^
5) Logical Operators
   Example: && || !

1) Comparison Operators:
=========================
* compare two operands and yield a Boolean value.
* allow values of the same data type for comparisons
    Return True if it matches the condition. 
    Return False if doesn't matches the condition.
# Common Comparisons:
  * Does one string match another?
  * Are two numbers the same?
  * Is one number greater than another

In comparison Operators are,
equal ==
not equal !=
less than <
less than or equal to <=
greater than >
greater than or equal to >=

Example program:
# main.go
package main
import ("fmt")
func main(){
    var city string  = "Chennai"
    var city2 string = "Madras"
    fmt.Printf(city == city2)
}
// When we run the program with,
go run main.go
false
# Note: return True when the values are equal.

package main
import "fmt"

func main() {
	var a int = 10
	var b int = 12
	fmt.Printf("%v", a >= b)
}
# Note: The Output will be false because a (10) is not greater than or equal to b (12)

package main
import ("fmt")
func main(){
	var a string = "100"
	var b string = "90"
	fmt.Printf("%v", a <= b)
}

go run main.go 
true

## Arithmetic Operators:-
=========================
It Used to perform common arithmetic operations,such as addition, subtraction, multiplication, etc..
 * Common comparisons:-
     * Does the sum of two numbers equal a particular value?
     * Is the difference between two numbers lesser than a particular value?

## Following are the common arithmetic operators:-
==================================================
addition (+)         // adds the left and right operand
subtraction (-)     // subtract the right operand from the left operand
multiplication (*) // multiplies both operands.
division (/)      // returns the quotient when left operand is divided by right operand
modulus (%)      // returns the remainder when left operand is divided by right operand.
increment (++)  // Unary operator.It increments the value of the operand by one.
decrement (--) // Unary operator. It decrements the value of the operand by one.

### Logical Operators:-
========================
It used to determine the logic between variables and values.
   * Common logical comparisons:
        * Are two variables both true?
        * Does either of two expressions evaluate to true?
