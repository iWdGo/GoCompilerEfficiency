package formparse

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("readkey").Parse(""))

func FormParseGlobal(w http.ResponseWriter, r *http.Request) (v string) {
	if err := tmpl.Execute(w, r.ParseForm()); err != nil {
		log.Fatalf("parse: %v", err)
	}
	return r.FormValue("value")
}

func FormParse(w http.ResponseWriter, r *http.Request) (v string) {
	// template is rebuild everytime
	if err := template.Must(template.New("readkey").Parse("")).Execute(w, r.ParseForm()); err != nil {
		log.Fatalf("parse: %v", err)
	}
	return r.FormValue("value")
}
