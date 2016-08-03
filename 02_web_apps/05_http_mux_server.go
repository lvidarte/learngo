package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("/home"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
