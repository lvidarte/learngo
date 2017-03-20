package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at",
		p.Address.Number,
		p.Address.Street,
		p.Address.City,
		p.Address.State,
		p.Address.Zip)
}

// Pseudo Subtyping
type Citizen struct {
	Country string
	Person  // This way we "extend" Person
}

func (c *Citizen) Nationality() {
	// We call .Name rathen than c.Person.Name
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

// Overwriting methods
func (c *Citizen) Talk() {
	fmt.Println("Hello, my name is", c.Name, "and I'm from", c.Country)
}

func main() {
	c := Citizen{
		Country: "Canada",
		Person: Person{
			Name: "Steve",
			Address: Address{
				Number: "13",
				Street: "Main",
				City:   "Gotham",
				State:  "NY",
				Zip:    "01313",
			},
		},
	}

	c.Talk()        // <- Notice both are accessible
	c.Person.Talk() // <- Notice both are accessible
	c.Location()
	c.Nationality()
}
