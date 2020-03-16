package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT")
}

func main() {
	http.HandleFunc("/", index)      // page 3000/
	http.HandleFunc("/about", about) // page 3000/about
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
