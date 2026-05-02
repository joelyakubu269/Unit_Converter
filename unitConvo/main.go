package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"strconv"
)
func handleLength(w http.ResponseWriter, r *http.Request) {
	n,from,to,err := parseValue(r)
	if err!= nil {
		http.Error(w,"only post allowed",http.StatusBadRequest)
		return
	}

	var result float64

	switch {
	case from == "kilometres" && to == "metres":
		result = n * 1000
	case from == "metres" && to == "kilometres":
		result = n / 1000
	case from == "metres" && to == "centimetres":
		result = n * 100
	case from == "centimetres" && to == "metres":
		result = n / 100
	case from == "kilometres" && to == "centimetres":
		result = n * 100000
	case from == "centimetres" && to == "kilometres":
		result = n / 100000
	default:
		http.Error(w, "invalid conversion", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"result": result,
	})
}

func handleWeight(w http.ResponseWriter, r *http.Request) {
	n, from, to, err := parseValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result float64

	switch {
	
	case from == "kilograms" && to == "grams":
		result = n * 1000
	case from == "grams" && to == "kilograms":
		result = n / 1000

	case from == "kilograms" && to == "pounds":
		result = n * 2.20462
	case from == "pounds" && to == "kilograms":
		result = n * 0.453592

	
	case from == "grams" && to == "pounds":
		result = n * 0.00220462
	case from == "pounds" && to == "grams":
		result = n / 0.00220462

	
	case from == to:
		result = n

	default:
		http.Error(w, "invalid weight conversion", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"result": result,
	})
}
func handleTemp(w http.ResponseWriter, r *http.Request) {
	n, from, to, err := parseValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result float64

	switch {
	
	case from == "celsius" && to == "kelvin":
		result = n + 273.15
	case from == "kelvin" && to == "celsius":
		result = n - 273.15

	
	case from == "celsius" && to == "fahrenheit":
		result = (n * 9/5) + 32
	case from == "fahrenheit" && to == "celsius":
		result = (n - 32) * 5/9

	
	case from == "kelvin" && to == "fahrenheit":
		result = (n-273.15)*9/5 + 32
	case from == "fahrenheit" && to == "kelvin":
		result = (n-32)*5/9 + 273.15

	
	case from == to:
		result = n

	default:
		http.Error(w, "invalid temperature conversion", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"result": result,
	})
}
func parseValue(r *http.Request) (float64, string, string, error){
if r.Method != http.MethodPost {
		return 0, "", "", fmt.Errorf("only POST allowed")
	}

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
			return 0, "", "", fmt.Errorf("invalid number: %v", err)
	}
 return n,from,to,nil
}

func main() {
	http.HandleFunc("/index.html", handleLength)
	http.HandleFunc("/weight.html", handleWeight)
	http.HandleFunc("/temperature.html", handleTemp)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080",nil)

}
