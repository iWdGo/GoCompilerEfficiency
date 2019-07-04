package writebuffer

import (
	"testing"
)

func TestFmtPrintBuffer(t *testing.T) {
	if i := bufferFmt(newBuffer()).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestBufferWrite(t *testing.T) {
	if i := bufferWrite(newBuffer()).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestBufferWriteTo(t *testing.T) {
	if i := bufferWriteTo(newBuffer()).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestIoCopyBuffer(t *testing.T) {
	if i := ioCopyBuffer(newBuffer()).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFmtPrintBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if i := bufferFmt(newBuffer()).Len(); i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkWriteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if i := bufferWrite(newBuffer()).Len(); i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkWriteToBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if i := bufferWriteTo(newBuffer()).Len(); i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkIoCopyBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if i := ioCopyBuffer(newBuffer()).Len(); i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}
