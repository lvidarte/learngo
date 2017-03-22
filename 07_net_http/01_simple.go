package main

import (
	"fmt"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello")
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe("localhost:8000", nil)
}
