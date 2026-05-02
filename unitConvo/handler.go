package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}
	n, from, to, convType, err := parseForm(r)
	if err != nil {
		tmpl.Execute(w, PageData{Error: err.Error()})
		return

	}
	result, err := convert(n,from,to,convType)
	if err!= nil {
		tmpl.Execute(w,PageData{Error: err.Error()})
		return
	}
		tmpl.Execute(w, PageData{Result: result})
}
