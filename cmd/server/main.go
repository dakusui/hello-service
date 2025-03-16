package main

import (
	"fmt"
	"log"
	"net/http"

	"hello-service/internal/config"
	"hello-service/internal/handler"
	"hello-service/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mux := http.NewServeMux()
	greetWithMiddleware := middleware.LogMiddleware(http.HandlerFunc(handler.GreetHandler))
	mux.Handle("/greet", greetWithMiddleware)

	port := cfg.Server.Port
	fmt.Printf("Starting server at :%d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
