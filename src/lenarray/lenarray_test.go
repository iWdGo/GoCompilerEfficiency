package lenarray

import "testing"

func TestLenOfArray(t *testing.T) {
	want := 999
	if got := lenOfArray(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestLenOfArrayVar(t *testing.T) {
	want := 999
	if got := lenOfArrayVar(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkLenOfArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lenOfArray()
	}
}
func BenchmarkLenOfArrayVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lenOfArrayVar()
	}
}
