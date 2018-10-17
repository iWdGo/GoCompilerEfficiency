package forbreak

import "testing"

/*
goos: windows
goarch: amd64
BenchmarkForGlobalVar-4           500000              2390 ns/op
BenchmarkForLocalVar-4           3000000               534 ns/op
*/
func TestForGlobalVar(t *testing.T) {
	want := 499500 // sum of 1 to 999
	if got := ForGlobalVar(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestForLocalVar(t *testing.T) {
	want := 499500 // sum of 1 to 999
	if got := ForLocalVar(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkForGlobalVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ForGlobalVar()
	}
}
func BenchmarkForLocalVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ForLocalVar()
	}
}
