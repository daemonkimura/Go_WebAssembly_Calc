package main

import (
	"net/http"
)

func main() {
	// handle static assets
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
