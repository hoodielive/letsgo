package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Demiurgos Aeon</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/heath_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
