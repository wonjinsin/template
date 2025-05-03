package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// get env variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Secret is this %s\n", secretKey)
	})
	http.ListenAndServe(":"+port, nil)
}
