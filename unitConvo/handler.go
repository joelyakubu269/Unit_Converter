package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("*.html"))

func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	n, from, to, err := parseForm(r)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return

	}
	result, err := convert(n, from, to, r.URL.Path)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return
	}
	tmpl.Execute(w, PageData{Result: result})
}
