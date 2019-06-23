package formparse

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRequest(s string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_ = r.ParseForm()
	r.Form.Set("value", s)
	return w, r
}

func TestFormParse(t *testing.T) {
	want := t.Name()
	if got := FormParse(newRequest(want)); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestFormParseGlobal(t *testing.T) {
	want := t.Name()
	if got := FormParseGlobal(newRequest(want)); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestFormParseImplicit(t *testing.T) {
	want := t.Name()
	if got := FormParseImplicit(newRequest(t.Name())); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFormParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		want := "example"
		if got := FormParse(newRequest(want)); got != want {
			b.Fatalf("got %s, want %s", got, want)
		}
	}
}

func BenchmarkFormParseGlobal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		want := "example"
		if got := FormParseGlobal(newRequest(want)); got != want {
			b.Fatalf("got %s, want %s", got, want)
		}
	}
}

func BenchmarkFormParseImplicit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		want := "example"
		if got := FormParseImplicit(newRequest(want)); got != want {
			b.Fatalf("got %s, want %s", got, want)
		}
	}
}
