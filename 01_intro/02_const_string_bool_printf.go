package main

import "fmt"

func main() {

	// A constant is a variable with a value that can't be changed
	const pi float64 = 3.1416

	// You can declare multiple variables like this
	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB, pi)

	// Strings are a serie of characters between " or `
	var myName string = "Leo Vidarte"

	// Get string length
	fmt.Println(len(myName))

	// You can combine strings with +
	fmt.Println(myName + " is a robot")

	// Booleans can be either true or false
	var isOver bool = true
	fmt.Println(isOver)

	// Printf is used for format printing (%f is for floats)
	fmt.Printf("%f \n", pi)

	// You can also define the decimal precision of a float
	fmt.Printf("%.3f \n", pi)

	// %c prints the character associated with a keycode
	fmt.Printf("%c \n", 44)

	// %x prints in hexcode
	fmt.Printf("%x \n", 17)

	// %e prints in scientific notation
	fmt.Printf("%e \n", pi)
}
