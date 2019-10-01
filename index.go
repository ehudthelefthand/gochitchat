package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"view/layout.html",
		"view/nav.html",
		"view/content.html",
	}
	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", struct{}{})
}
