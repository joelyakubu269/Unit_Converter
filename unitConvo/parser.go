package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseForm(r *http.Request) (float64, string, string, string, error) {

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
