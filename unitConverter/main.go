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
	tmpl := template.Must(template.ParseFiles("index.html"))

	if r.Method == http.MethodPost {
		value := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var result float64

		if from == "Celsius" && to == "Fahrenheit" {
			result = (num * 9 / 5) + 32
		} else if from == "Celsius" && to == "Kelvin" {
			result = num + 273.15
		} else if from == "Fahrenheit" && to == "Celsius" {
			result = (num - 32) * 5 / 9
		} else if from == "Fahrenheit" && to == "Kelvin" {
			result = (num-32)*5/9 + 273.15
		} else if from == "Kelvin" && to == "Celsius" {
			result = num - 273.15
		} else if from == "Kelvin" && to == "Fahrenheit" {
			result = (num-273.15)*9/5 + 32
		}

		data := Result{
			Input:  value + " " + from,
			Output: fmt.Sprintf("%.4f %s", result, to),
		}

		tmpl.Execute(w, data)
		return
	}

	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
