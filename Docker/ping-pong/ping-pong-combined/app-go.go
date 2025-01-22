package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from Go")
	})
	fmt.Println("Go server is running on port 3002...")
	http.ListenAndServe(":3002", nil)
}
