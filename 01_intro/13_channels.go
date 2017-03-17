package main

import "fmt"
import "time"

func get(messages chan string, i int) {
	fmt.Println(fmt.Sprintf("get START (%d)", i))
	fmt.Println(fmt.Sprintf("get DONE (%d): %s", i, <-messages))
}

func main() {
	messages := make(chan string)

	go get(messages, 0)
	go get(messages, 1)
	go get(messages, 2)
	go get(messages, 3)
	go get(messages, 4)
	go get(messages, 5)
	go get(messages, 6)

	for i := 0; i < 10; i++ {
		go func(messages chan string, message string, i int) {
			fmt.Println(fmt.Sprintf("\tput START (%d): %s", i, message))
			messages <- message
			fmt.Println(fmt.Sprintf("\tput DONE (%d): %s", i, message))
		}(messages, fmt.Sprintf("hello %d", i), i)
	}

	go get(messages, 7)
	go get(messages, 8)
	go get(messages, 9)

	time.Sleep(10000 * time.Millisecond)
}
