package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func api() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
