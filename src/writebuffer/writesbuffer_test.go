package writebuffer

import (
	"fmt"
	"testing"
)

func TestFmtWriteBuffer(t *testing.T) {
	if i := bufferFmt(newBuffer()).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestIoWriteBuffer(t *testing.T) {
	if i := bufferWriteTo(newBuffer()).Len(); i != len(s)*len(s) {
		fmt.Println(bufferWriteTo(newBuffer()))
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFmtWriteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := bufferFmt(newBuffer()).Len()
		if i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkIoWriteBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := bufferWriteTo(newBuffer()).Len()
		if i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}
