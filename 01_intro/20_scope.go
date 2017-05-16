package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		x := i
		fmt.Println(x)
	}
	fmt.Println(i)
	fmt.Println(x)
}
