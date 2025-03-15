package main

import (
  "fmt"
  "log"
  "net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Query().Get("name")
  if name == "" {
    name = "World"
  }

  message := fmt.Sprintf("Hello, %s!", name)

  fmt.Fprintln(w, message)
}

func main() {
  http.HandleFunc("/greet", greetHandler)

  fmt.Println("Starting server at :8080")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal(err)
  }
}
  