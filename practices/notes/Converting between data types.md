## Converting between data types:

```bash
   Type Casting:
     * The process is Converting one data type to another is known as Type Casting.
     * Data types can be converted to other data types, but this does not guarantee
     that the value will remain intact.

  Example for converting integer into float:-

  # main.go
package main

import (
	"fmt"
)

func main() {
	var i int = 40
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
}

# When you run the program with
   go run main.go
# The output will be
   40.00

  Example for converting float into integer:-
# main.go
package main

import "fmt"

func main() {
	var f float64 = 100.2343
	var i int = int(f)
	fmt.Printf("%d\n", i)
}

# When you run the program with
   go run main.go
# The output will be
   40.00

2) strconv package (string convertion package):

Itoa() - stands for integer to ASCII
  * converts integer to string
  * returns one value - string formed with the given integer.

Atoi() - stand for ASCII to Integer
  * converst string to integer
  * returns two values - the corresponding integer, error (if any)

Example for converting integer to string:

# main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q", s)
}

# When you run the program with
   go run main.go
# The output will be
   "42"

Example for Converting string to integer:

# main.go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Converting String to Integer using 'Atoi()' function..\n")
	var s string = "2000"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", err, err)
}

# When you run the program with
   go run main.go
# The output will be
   Converting String to Integer using 'Atoi()' function..
   2000 int
   <nil> <nil>



