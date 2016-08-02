package main

import "fmt"

func main() {

	// If statement

	yourAge := 18

	// You can use else if perform different actions,
	// but once a match is reached the rest of the conditions
	// aren't checked

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can have fun")
	}

	// Switch statements are used when you have limited options

	switch yourAge {
	case 16:
		fmt.Println("Go drive")
	case 18:
		fmt.Println("Go vote")
	default:
		fmt.Println("Go have fun")
	}
}
