package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseForm(r *http.Request) (float64, string, string, string, error) {
	if err := r.ParseForm(); err != nil { // r.ParseForm() is used to get all values in name
		return 0, "", "", "", fmt.Errorf("error getting values") // is using this error wrapper okay
	}
	value := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")
	convType := r.FormValue("type")
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, "", "", "", fmt.Errorf("invalid Number")
	}
	return n, from, to, convType, err
}
