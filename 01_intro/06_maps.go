package main

import (
	"fmt"
)

func main() {

	// A Map is a collection of key value pairs
	// Created with varName := make(map[keyType] valueType)

	myMap := make(map[string]int)

	myMap["one"] = 1

	fmt.Println("key one", myMap["one"])

	// Get the number of items in the map

	fmt.Println("len", len(myMap))

	// The size changes when a new item is added

	myMap["two"] = 2
	fmt.Println("len after resize", len(myMap))

	// We can delete by passing the key to delete

	delete(myMap, "one")
	fmt.Println("len after delete", len(myMap))

	// Print the whole map
	myMap["three"] = 3
	fmt.Println(myMap)

}
