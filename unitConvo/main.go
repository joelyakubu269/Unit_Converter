package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/length", handleConvert)
	http.HandleFunc("/weight", handleConvert)
	http.HandleFunc("/temperature", handleConvert)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080", nil)

}
