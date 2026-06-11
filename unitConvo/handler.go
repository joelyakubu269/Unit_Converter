package main

import (
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseGlob("*.html"))

func handleConvert(w http.ResponseWriter, r *http.Request) {
	convType := strings.TrimPrefix(r.URL.Path, "/")

	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, convType+".html", nil) // did convtype so i can know which page to return
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Input, err := parseForm(r)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return

	}
	switch convType {
	case "length":
		result, err := convertGeneric(Input.Value, Input.From, Input.To, lengthUnits)
		if err != nil {
			tmpl.Execute(w, PageData{Error: err.Error()})
			return
		}
		tmpl.Execute(w, PageData{Result: result})
	case "weight":
		result, err := convertGeneric(Input.Value, Input.From, Input.To, weightUnits)
		if err != nil {
			tmpl.Execute(w, PageData{Error: err.Error()})
			return
		}
		tmpl.Execute(w, PageData{Result: result})
	case "temperature":
		result, err := convertGeneric(Input.Value, Input.From, Input.To, tempUnits)
		if err != nil {
			tmpl.Execute(w, PageData{Error: err.Error()})
			return
		}
		tmpl.Execute(w, PageData{Result: result})
	}
}

func handleExecute(float64, error)
