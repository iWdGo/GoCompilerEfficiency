package bitorif

import (
	"os"
	"testing"
)

const null = "nuG"

func TestIsNulNameBit(t *testing.T) {
	tests := []struct {
		n string
		b bool
	}{
		{"nul", true},
		{"NUL", true},
		{"Nul", true},
		{"NUg", false},
		{"NUG", false},
	}
	for _, s := range tests {
		if g := isNulNameBit(s.n); g != s.b {
			t.Fatalf("%s is not 'nul' (whatever case) but was not detected", s.n)
		}
	}
}

func TestIsNulNameIf(t *testing.T) {
	if isNulNameIf(null) {
		t.Fatalf("%s is not nul or NUL but was not detected", null)
	}
}

func TestIsNulNameLoop(t *testing.T) {
	if !isNulNameLoop(os.DevNull) {
		t.Fatalf("os.DevNull was not detected")
	}
	if isNulNameLoop(os.DevNull[:len(os.DevNull)-1] + ".") {
		t.Fatalf("%s is not nul or NUL but was not detected", os.DevNull)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkBitMasking(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isNulNameBit(null)
	}
}
func BenchmarkSecondIf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isNulNameIf(null)
	}
}

func BenchmarkLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isNulNameLoop(os.DevNull)
	}
}
