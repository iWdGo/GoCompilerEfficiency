package firstitem

import "testing"

var want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestInitArrayMake(t *testing.T) {
	if got := InitArrayMake(); got[len(got)-1] != want[len(want)-1] {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestInitArrayStatic(t *testing.T) {
	if got := InitArrayStatic(); got[len(got)-1] != want[len(want)-1] {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestInitArrayAppend(t *testing.T) {
	if got := InitArrayAppend(); got[len(got)-1] != want[len(want)-1] {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestInitArrayAppendStatic(t *testing.T) {
	if got := InitArrayAppendStatic(); got[len(got)-1] != want[len(want)-1] {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkInitArrayMake(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitArrayMake()
	}
}
func BenchmarkInitArrayStatic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitArrayStatic()
	}
}
func BenchmarkInitArrayAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitArrayAppend()
	}
}
func BenchmarkInitArrayAppendStatic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitArrayAppendStatic()
	}
}
