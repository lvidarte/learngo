package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var age int64 = 40
	var price = 1.63434
	randNum := 1

	fmt.Println(age, price, randNum)

	var numOne = 1.000
	var num99 = .999

	fmt.Println(numOne - num99)

	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)
}
