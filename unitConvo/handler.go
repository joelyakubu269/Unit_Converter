package main

import (
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseGlob("*.html"))

//var units map[string]Unit // avoid using shared variables because of race condition

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
	unit := getUnits(convType)
	if unit == nil {
		http.NotFound(w, r)
		return
	}
	result, err := convertGeneric(Input.Value, Input.From, Input.To, unit)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return
	}
	tmpl.Execute(w, PageData{Result: result})

}
