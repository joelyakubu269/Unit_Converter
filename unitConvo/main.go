package main

import (
	"fmt"
	"net/http"

	"strconv"
)

func parseForm(r *http.Request) (Input, error) {

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")
	if value == "" || from == "" || to == "" {
		return Input{}, fmt.Errorf("required fields are missing")
	}
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return Input{}, fmt.Errorf("invalid Number %v", err)
	}
	return Input{
		Value: n,
		From:  from,
		To:    to,
	}, nil
}

func main() {

	http.HandleFunc("/length", handleLength)
	http.HandleFunc("/weight", handleWeight)
	http.HandleFunc("/temperature", handleTemp)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080", nil)

}
