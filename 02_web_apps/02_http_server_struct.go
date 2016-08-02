package main

import (
	"net/http"
)

type person struct {
	name string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + p.name))
}

func main() {
	personOne := &person{name: "Leo"}

	http.ListenAndServe(":9001", personOne)
}
