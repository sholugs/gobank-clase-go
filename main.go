package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the PORT from the environment variables, if not defined, use 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the HTTP server
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
