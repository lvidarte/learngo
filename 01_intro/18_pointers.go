package main

import (
	"fmt"
	"reflect"
)

func describe(i interface{}, pi interface{}) {
	fmt.Println(
		i, reflect.TypeOf(i),
		pi, reflect.TypeOf(pi),
	)
}

func main() {

	var (
		x = 2
		y = false
		z = 0.1
	)

	var (
		px *int     = &x
		py *bool    = &y
		pz *float64 = &z
	)

	describe(x, px)
	describe(y, py)
	describe(z, pz)
}
