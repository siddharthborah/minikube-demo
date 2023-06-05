package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Printf("Received request\n")
	fmt.Fprintf(w, "<a href=\"https://www.w3schools.com\">Visit W3Schools</a>")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is the error")
}
