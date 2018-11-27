package isvalid

import "testing"

var (
	want  = []int{-1, 2, 0, 8, 10} // got == []int{1, 8}
	wantL = 2                      // valid values
)

func TestCheckValueFuncVar(t *testing.T) {
	if got := CheckValueFuncVar(want); len(got) != wantL {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestCheckValueIfVar(t *testing.T) {
	if got := CheckValueIfVar(want); len(got) != wantL {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestCheckValueForFunc(t *testing.T) {
	if got := CheckValueForFunc(want); len(got) != wantL {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
func TestCheckValueForIf(t *testing.T) {
	if got := CheckValueForIf(want); len(got) != wantL {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkCheckValueFuncVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckValueFuncVar(want)
	}
}
func BenchmarkCheckValueIfVar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckValueIfVar(want)
	}
}
func BenchmarkCheckValueForFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckValueForFunc(want)
	}
}
func BenchmarkCheckValueForIf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckValueForIf(want)
	}
}
