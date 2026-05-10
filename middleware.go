package main

import (
	"fmt"
	"net/http"
)

// midddleware
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
		next(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func main() {
	http.HandleFunc("/", middleware(home))

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
