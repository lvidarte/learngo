package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler logic into a closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path + " 200")
		fmt.Fprintf(w, message)
	})
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(r.URL.Path[1:])
	if err == nil {
		log.Println(r.URL.Path + " 200")
		w.Write(data)
	} else {
		w.WriteHeader(404)
		log.Println(r.URL.Path + " 404")
		w.Write([]byte("404 - Not found"))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", fileHandler)

	mux.Handle("/hello", messageHandler("Welcome to Go Web Development"))
	mux.Handle("/net/http", messageHandler("net/http is awesome"))

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
