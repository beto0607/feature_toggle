package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("templates/index.html"))
	templ.Execute(w, nil)
}
