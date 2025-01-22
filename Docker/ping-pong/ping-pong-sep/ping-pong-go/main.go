package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ping" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "pong")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("Server starting on port 8081...")
	http.ListenAndServe(":8081", nil)
}
