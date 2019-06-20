package readbyte

// Counting bytes until EOF
import (
	"testing"
)

func TestBufferRead(t *testing.T) {
	b := fillBuffer(times)
	if got := ReadAllBytesRead(b); got != want {
		t.Fatalf("buffer read failed: want %d, got %d\n", want, got)
	}
}

func TestBufferReadByte(t *testing.T) {
	b := fillBuffer(times)
	if got := ReadAllBytesReadByte(b); got != want {
		t.Fatalf("buffer read failed: want %d, got %d\n", want, got)
	}
}

/*
func TestBufferReadDelim(t *testing.T) {
}
*/
// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkBufferRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadAllBytesRead(fillBuffer(times))
	}
}
func BenchmarkRBufferReadByte(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadAllBytesReadByte(fillBuffer(times))
	}
}

/*
func BenchmarkReturnBufferBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReturnBufferBytes(times)
	}
}
*/
