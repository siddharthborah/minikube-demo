package main

import (
	"fmt"
	"log"
	"minikube-demo/handlers"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
