package returnbuffer

import (
	"testing"
)

const (
	times = 20
	want  = len(filling)*times + 9 + 2*(times-9) // 8 one digit + 11 2-digits (11 to 20)
)

func TestReturnBuffer(t *testing.T) {
	tb := ReturnBuffer(times)
	if got := tb.Len(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestReturnBufferString(t *testing.T) {
	tb := ReturnBufferString(times)
	if got := len(tb); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestReturnBufferBytes(t *testing.T) {
	tb := ReturnBufferBytes(times)
	if got := len(tb); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkReturnBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReturnBuffer(times)
	}
}
func BenchmarkReturnBufferString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReturnBufferString(times)
	}
}
func BenchmarkReturnBufferBytes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReturnBufferBytes(times)
	}
}
