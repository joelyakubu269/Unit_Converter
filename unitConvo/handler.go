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
		tmpl.ExecuteTemplate(w, convType+".html", nil)
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
	result, err := convert(Input.Value, Input.From, Input.To, convType)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return
	}
	tmpl.Execute(w, PageData{Result: result})
}
