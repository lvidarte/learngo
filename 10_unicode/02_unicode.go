package main

import (
	"fmt"
)

func main() {
	const uni = `π`

	fmt.Printf("len %d\n", len(uni))

	fmt.Printf("%s\n", uni)

	fmt.Printf("%+q\n", uni)

	fmt.Printf("%x\n", uni)

	for i := 0; i < len(uni); i++ {
		fmt.Printf("%x ", uni[i])
	}
}
