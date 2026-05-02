package main

import (
	"fmt"
	"net/http"
)

func parseForm(r *http.Request) (float64, string, string, string, error) {
	if r.Method != http.MethodPost {
		return 0,"","","",fmt.Errorf("only post method allowed")
	}
}
