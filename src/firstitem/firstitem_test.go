package firstitem

import "testing"

const (
	loops = 1000
	want  = (loops - 1) + (loops - 2)
)

func TestFirstItemBefore(t *testing.T) {
	if got := FirstItemBefore(loops); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestFirstItemIfOnIndex(t *testing.T) {
	if got := FirstItemIfOnIndex(loops); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestFirstItemIfOnArray(t *testing.T) {
	if got := FirstItemIfOnArray(loops); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestFirstItemBeforeUnfair(t *testing.T) {
	if got := FirstItemBeforeUnfair(loops); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFirstItemBefore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemBefore(loops)
	}
}
func BenchmarkFirstItemIfOnIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemIfOnArray(loops)
	}
}
func BenchmarkFirstItemIfOnArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemIfOnIndex(loops)
	}
}
func BenchmarkFirstItemBeforeUnfair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemBeforeUnfair(loops)
	}
}
