package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go away!")
}

func handleRequests() {
	http.HandleFunc("/", hello)
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}
