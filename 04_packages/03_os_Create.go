package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
	}
	defer file.Close()

	file.WriteString("test")
}
