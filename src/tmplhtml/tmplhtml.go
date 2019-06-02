package tmplhtml

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

const docPage = "efficiency"

var (
	tmpl = `<a href="http://golang.org/{{.V}}">documentation</a>`
	// 10x faster than parsing every time as elsewhere
	tmplT = template.Must(template.New("mytmpl").Parse(tmpl))
)

type data = struct {
	V string
}

// The test only executes the template with one variable value.
func hTMLOutputTmpl(d *data) *bytes.Buffer {
	/* Passing the string or the struct does not change much the performance
	d := new(data)
	d.V = s

	*/
	b := new(bytes.Buffer)
	if err := tmplT.Execute(b, d); err != nil {
		log.Fatalf("parse: %v", err)
	}
	return b
}

func hTMLOutputBuffer(s string) *bytes.Buffer {
	b := new(bytes.Buffer)
	if _, err := fmt.Fprint(b, `<a href="http://golang.org/`+s+`">documentation</a>`); err != nil {
		log.Fatalf("fprint: %v", err)
	}
	return b
}
