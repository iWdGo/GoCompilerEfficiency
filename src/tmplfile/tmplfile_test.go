package tmplfile

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoadTemplateFile(t *testing.T) {
	w := httptest.NewRecorder()
	if !strings.Contains(fmt.Sprint(loadTemplateFile(w)), ValueToFind) {
		t.Fatalf("%s not found", ValueToFind)
	}
}

func TestLoadTemplateFileOnce(t *testing.T) {
	w := httptest.NewRecorder()
	if !strings.Contains(fmt.Sprint(loadTemplateFileOnce(w)), ValueToFind) {
		t.Fatalf("%s not found", ValueToFind)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkLoadTemplateFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		w := httptest.NewRecorder()
		if !strings.Contains(fmt.Sprint(loadTemplateFile(w)), ValueToFind) {
			b.Fatalf("%s not found", ValueToFind)
		}
	}
}
func BenchmarkLoadTemplateFileOnce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		w := httptest.NewRecorder()
		if !strings.Contains(fmt.Sprint(loadTemplateFileOnce(w)), ValueToFind) {
			b.Fatalf("%s not found", ValueToFind)
		}
	}
}
