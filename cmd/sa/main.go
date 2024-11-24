package main

import (
	"log"
	"net/http"

	"github.com/tmhdgsn/sa/web"
)

func main() {
	// Register the static file handler
	http.Handle("/", web.Handler())

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
