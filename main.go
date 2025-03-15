package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Wire handler with middleware
	greetWithMiddleware := logMiddleware(http.HandlerFunc(greetHandler))

	http.Handle("/greet", greetWithMiddleware)

	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
