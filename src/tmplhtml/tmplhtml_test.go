package tmplhtml

import "testing"

const (
	page      = "efficiency"
	linkValue = `<a href="http://golang.org/efficiency">documentation</a>`
)

func TestHTMLOutputTmpl(t *testing.T) {
	p := new(data)
	p.V = page
	for i := 0; i < 5; i++ {
		if got := hTMLOutputTmpl(p); got.String() != linkValue {
			t.Fatalf("want %s, got %s\n", linkValue, got)
		}
	}
}

func TestHTMLOutputString(t *testing.T) {
	// For fairness but no change
	var s string
	s = page
	for i := 0; i < 5; i++ {
		if got := hTMLOutputBuffer(s); got.String() != linkValue {
			t.Fatalf("want %s, got %s\n", linkValue, got)
		}
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkHTMLOutputTmpl(b *testing.B) {
	p := new(data)
	p.V = page
	for n := 0; n < b.N; n++ {
		hTMLOutputTmpl(p)
	}
}
func BenchmarkHTMLOutputBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hTMLOutputBuffer(page)
	}
}
