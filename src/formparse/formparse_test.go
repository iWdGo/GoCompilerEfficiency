package formparse

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const want = "example"

func TestFormParse(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	r.Form.Set("value", "example")

	if got := FormParse(w, r); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestFormParseGlobal(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	r.Form.Set("value", "example")

	if got := FormParseGlobal(w, r); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFormParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "/", http.NoBody)
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.ParseForm()
		r.Form.Set("value", "example")

		if got := FormParse(w, r); got != want {
			b.Fatalf("got %s, want %s", got, want)
		}
	}
}
func BenchmarkFormParseGlobal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "/", http.NoBody)
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.ParseForm()
		r.Form.Set("value", "example")

		if got := FormParseGlobal(w, r); got != want {
			b.Fatalf("got %s, want %s", got, want)
		}
	}
}
