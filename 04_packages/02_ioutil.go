package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bs, err := ioutil.ReadFile("tet.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
		return
	}
	str := string(bs)
	fmt.Println(str)
}
