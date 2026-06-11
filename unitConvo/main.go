package main

import (
	"fmt"
	"net/http"

	"strconv"
)

//	func home(w http.ResponseWriter, r *http.Request) {
//		if r.Method == http.MethodGet {
//			http.ServeFile(w, r, "index.html")
//			return
//		}
//	}
func handleLength(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r)
}

func handleWeight(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r)
}
func handleTemp(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r)
}

func parseForm(r *http.Request) (float64, string, string, error) {

	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")
	//convType := r.FormValue("type")
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, "", "", fmt.Errorf("invalid Number")
	}
	return n, from, to, err
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/length", handleLength)
	http.HandleFunc("/weight", handleWeight)
	http.HandleFunc("/temperature", handleTemp)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080", nil)

}
