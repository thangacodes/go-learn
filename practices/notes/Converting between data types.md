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
