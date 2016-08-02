package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	data, err := ioutil.ReadFile(path)

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		}

		log.Print("200 " + path)
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		log.Print("404 " + path)
		w.WriteHeader(404)
		w.Write([]byte("404 - Not found " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":9000", nil)
}
