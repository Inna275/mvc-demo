package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
