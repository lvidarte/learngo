package main

import (
	"fmt"
	//"math"
)

type Person struct {
	Name string
}

func (p *Person) Talk() (s string) {
	s = "Hi, my name is " + p.Name
	return
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Person.Name = "Leo"
	fmt.Println(a.Person.Talk())
}
