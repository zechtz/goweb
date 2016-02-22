package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", messageHandler("Welcome to Go Web Development"))
	mux.Handle("/", messageHandler("Welcome to the Homepage"))

	log.Println("Listening on port 8080....")
	http.ListenAndServe(":8080", mux)
}
