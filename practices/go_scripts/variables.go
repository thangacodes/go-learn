package main

import (
	"fmt"
)

func main() {
	name := "John Morris"
	dob := "12-07-1989"
	birth_place := "San Antonio, Texas"
	studied := "University of Texas at Austin"
	work := "Google, LLC"
	hobbies := "Playing Tennis, Rugby, Travelling"
	fmt.Println("My Name is:", name)
	fmt.Println("My DOB is :", dob)
	fmt.Println("My Birth place is:", birth_place)
	fmt.Println("I studied my B.S in: ", studied)
	fmt.Println("I'm currently working as a Software Engineer at:", work)
	fmt.Println("During my free time, I used to", hobbies)
}
