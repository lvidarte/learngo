package main

import (
	"fmt"
	"reflect"
)

/**
 * Basic types
 * -----------
 *
 * bool
 *
 * string
 *
 * int  int8  int16  int32  int64
 * uint uint8 uint16 uint32 uint64 uintptr
 *
 * byte // alias for uint8
 *
 * rune // alias for int32
 *      // represents a Unicode code point
 *
 * float32 float64
 *
 * complex64 complex128
 */

func main() {
	var x, y int8 = 1, 2
	z, w := 1, 2 // Short declaration

	fmt.Println(x, y, z, w)

	var s0 string = "hello"
	s1 := "world" // Short declaration

	fmt.Println(s0, s1)

	a, b, c, d := 1, true, "bye", 0.1 // Short declaration
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d))

	fmt.Println(a, b, c, d)

	var (
		n0, n1 int16
		m0, m1 float32 = 1.2, 3.14
		t0, t1 string  = "foo", "bar"
	)

	fmt.Println(n0, n1, m0, m1, t0, t1)
	fmt.Println(reflect.TypeOf(n0), reflect.TypeOf(m0), reflect.TypeOf(t0))
}
