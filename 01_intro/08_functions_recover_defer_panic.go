package main

import (
	"fmt"
)

func main() {

	listOfNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum :", addThemUp(listOfNums))

	// Get 2 values from a function

	num1, num2 := next2Values(5)

	fmt.Println(num1, num2)

	// Send an undefined number of values to a function
	// (Variadic Function)

	num3 := 3

	doubleNum := func() int {

		num3 *= 2
		return num3

	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// Calling a recursive function

	fmt.Println(factorial(3))

	// Defer executes a function after the inclosing function finishes
	// Defer can be used to keep functions together in a logical way
	// but at the same time execute on last as a clean up operation
	// Ex. Defer closing a file after we open it and perfom operations

	defer printTwo()
	printOne()

	// Use recover() to catch a division by 0 error

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	// We can catch our own errors and recover with panic & recover

	demPanic()
}

func addThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum

}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

// If a error occurs we can catch the error with recover and allow
// code to continue to execute

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2

	return solution
}

// Desmostrate how to call panic and handle it with recover

func demPanic() {
	defer func() {
		// If I didn't print the message nothing would show
		fmt.Println(recover())
	}()
	panic("PANIC")
}
