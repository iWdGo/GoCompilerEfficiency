package firstitem

import "testing"

const want = 1997 // 999 numbers and 998 "+"

func TestFirstItemBefore(t *testing.T) {
	if got := FirstItemBefore(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestFirstItemIf(t *testing.T) {
	if got := FirstItemIf(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestFirstItemBeforeUnfair(t *testing.T) {
	if got := FirstItemBeforeUnfair(); got != want {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFirstItemBefore(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemBefore()
	}
}
func BenchmarkFirstItemIf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemIf()
	}
}
func BenchmarkFirstItemBeforeUnfair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FirstItemBeforeUnfair()
	}
}
