package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle area =", getArea(rect))
	fmt.Println("Circle area =", getArea(circ))

}

// An interface defines a list of methods that a type must implement
// If that type implements those methods the proper method is executed
// even if the original is referred to with the interface name

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Nothing additional is required to implement an interface
// (there is no implements or extends keyword).
// It’s sufficient to merely have a method with the same name and
// signature.

func getArea(shape Shape) float64 {
	return shape.area()
}
