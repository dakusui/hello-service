package main

import (
	"fmt"
	"hello-service/internal/handler"
	"hello-service/internal/middleware"
	"log"
	"net/http"
)

func main() {
	// Wire handler with middleware
	greetWithMiddleware := middleware.LogMiddleware(http.HandlerFunc(handler.GreetHandler))

	http.Handle("/greet", greetWithMiddleware)

	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
