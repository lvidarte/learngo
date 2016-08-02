package main

import (
	"fmt"
)

func main() {

	// We pass the value of a variable to the function

	x := 0
	changeXVal(x)
	fmt.Println("x =", x)

	// If we pass a reference to the variable we can change the value
	// in a function

	changeXValNow(&x)
	fmt.Println("x =", x)

	// Get the address x points to with &

	fmt.Println("Memory Address for x=", &x)

	// We can also generate a pointer with new

	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)
}

func changeXVal(x int) {
	// Has no effect on the value of x in main()
	x = 2
}

func changeXValNow(x *int) {
	// Has no effect on the value of x in main()
	*x = 2
}

func changeYValNow(y *int) {
	*y = 100
}
