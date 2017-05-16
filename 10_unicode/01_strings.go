package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)
	fmt.Printf("% x\n", sample)
	fmt.Printf("%+q\n", sample)
}
