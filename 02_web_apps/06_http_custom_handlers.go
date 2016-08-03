package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	handler1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/", handler1)

	handler2 := &messageHandler{"net/http is awesome"}
	mux.Handle("/net/http", handler2)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
