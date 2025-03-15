package handler

import (
	"fmt"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, message)
}
