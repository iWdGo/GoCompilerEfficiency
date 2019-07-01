package writestring

import (
	"testing"
)

func TestFmtWriteString(t *testing.T) {
	if i := fmtWriteString(s).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestIoWriteString(t *testing.T) {
	if i := ioWriteString(s).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

func TestBufferWriteTo(t *testing.T) {
	if i := bytesWriteTo(s).Len(); i != len(s)*len(s) {
		t.Fatalf("got %d, want %d\n", i, len(s)*len(s))
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFmtWriteString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := fmtWriteString(s).Len()
		if i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkIoWriteString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := ioWriteString(s).Len()
		if i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}

func BenchmarkBufferWriteTo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := bytesWriteTo(s).Len()
		if i != len(s)*len(s) {
			b.Logf("got %d, want %d\n", i, len(s)*len(s))
		}
	}
}
