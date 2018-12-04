package emptycase

import "testing"

var (
	want = []int{10, 8, -1, 2, -1}
)

func TestCheckUselessCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		if got := CheckUselessCase(i); got != want[i] {
			t.Fatalf("want %d, got %d\n", want, got)
		}
	}
}
func TestCheckReturnDefault(t *testing.T) {
	for i := 0; i < 5; i++ {
		if got := CheckReturnDefault(i); got != want[i] {
			t.Fatalf("want %d, got %d\n", want, got)
		}
	}
}
func TestCheckNoUselessCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		if got := CheckNoUselessCase(i); got != want[i] {
			t.Fatalf("want %d, got %d\n", want, got)
		}
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkCheckUselessCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckUselessCase(2)
	}
}
func BenchmarkReturnDefault(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckReturnDefault(2)
	}
}
func BenchmarkNoUselessCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckNoUselessCase(2)
	}
}
