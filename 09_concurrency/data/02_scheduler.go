package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine")
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Inside a goroutine", i)
		}
	}()
	fmt.Println("Outside again")

	runtime.Gosched() // Yields to the scheduler

	fmt.Println("End")
}
