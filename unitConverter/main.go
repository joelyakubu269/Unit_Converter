package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Result struct {
	Input  string
	Output string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method == "http.MethodPost" {
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		var result float64
		if from == "meters" && to == "kilometres" {
			result = num / 1000
		} else if from == "kilometres" && to == "meters" {
			result = num * 1000
		} else if from == "meters" && to == "centimeters" {
			result = num * 100
		} else if from == "centimeters" && to == "meters" {
			result = num / 100
		}
		data := Result{
			Input:  value + " " + from,
			Output: fmt.Sprintf("%.4f %s", result, to),
		}
		tmpl.Execute(w, data)
		return

	}
	http.ServeFile(w, r, "index.html")
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
