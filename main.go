package main

import (
	"fmt"
	"go-web-app/pkg/handlers"

	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
