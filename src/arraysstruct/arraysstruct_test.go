package arraysstruct

import "testing"

const want = "this is only data to fill the array"

func TestMain(m *testing.M) {
	fill()
	m.Run()
}

func TestArrays(t *testing.T) {
	if got := Arrays(); got != want {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
func TestArrayStructs(t *testing.T) {
	if got := ArrayStructs(); got != want {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
func TestMapStructs(t *testing.T) {
	if got := MapStructs(); got != want {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkArrays(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Arrays()
	}
}
func BenchmarkArrayStructs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ArrayStructs()
	}
}
func BenchmarkMapStructs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapStructs()
	}
}
