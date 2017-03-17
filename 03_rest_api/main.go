package main

import (
  "fmt"
  "encoding/json"
  "net/http"
)

type Payload struct {
  Code int  `json:"code"`
  Data Data `json:"data"`
}

type Data struct {
  Fruits     Fruits     `json:"fruits"`
  Vegetables Vegetables `json:"vegetables"`
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
  response, err := getJsonResponse()
  if err != nil {
    panic(err)
  }
  fmt.Fprintf(w, string(response))
}

func main() {
  http.HandleFunc("/", serveRest)
  http.ListenAndServe("localhost:3000", nil)
}

func getJsonResponse() ([]byte, error) {
  fruits := make(map[string]int)
  fruits["apples"] = 12
  fruits["bananas"] = 2

  vegetables := make(map[string]int)
  vegetables["carrots"] = 2

  data := Data{fruits, vegetables}

  payload := Payload{200, data}

  return json.Marshal(payload)
}
