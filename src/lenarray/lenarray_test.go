package lenarray

import "testing"

/*
goos: windows
goarch: amd64
BenchmarkLenOfArray-4             200000              6147 ns/op
BenchmarkLenOfArrayVar-4          200000              6316 ns/op
PASS
*/
func TestLenOfArray(t *testing.T) {
	want := 999
	if got := LenOfArray(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestLenOfArrayVar(t *testing.T) {
	want := 999
	if got := LenOfArrayVar(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
/*
goos: windows
goarch: amd64
BenchmarkForGlobalVar-4           500000              2390 ns/op
BenchmarkForLocalVar-4           3000000               534 ns/op
*/
func BenchmarkLenOfArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LenOfArray()
	}
}
func BenchmarkLenOfArrayVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LenOfArrayVar()
	}
}
