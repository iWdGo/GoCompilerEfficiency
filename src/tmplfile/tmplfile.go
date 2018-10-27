package tmplfile

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const ValueToFind = "value-to-find"

var (
	t = loadTemplate("tmpl.html")

	data = struct {
		V string
	}{
		V: ValueToFind,
	}
)

func loadTemplate(n string) *template.Template {
	pattern := filepath.Join(n)
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		log.Fatalf("load template: %v", err)
	}
	return template.Must(tmpl, err)
}

func loadTemplateFile(w http.ResponseWriter) http.ResponseWriter {
	t := loadTemplate("tmpl.html")
	if err := t.Execute(w, data); err != nil {
		log.Fatalf("execution failed: %v", err)
	}
	return w
}

// Template is loaded once and passed to func
func loadTemplateFileOnce(w http.ResponseWriter) http.ResponseWriter {
	if err := t.Execute(w, data); err != nil {
		log.Fatalf("execution failed: %v", err)
	}
	return w
}
