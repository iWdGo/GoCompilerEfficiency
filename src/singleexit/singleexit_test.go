package singleexit

import "testing"

var testcases = []struct {
	want   bool
	search string
}{
	{true, "bingo"},
	{true, "long"},
	{false, "zorro"},
}

func TestFindString(t *testing.T) {
	for _, c := range testcases {
		if got := FindString(c.search); got != c.want {
			t.Fatalf("want %v, got %v\n", c.want, got)
		}
	}
}
func TestFindStringRange(t *testing.T) {
	for _, c := range testcases {
		if got := FindStringRange(c.search); got != c.want {
			t.Fatalf("want %v, got %v\n", c.want, got)
		}
	}
}

// TODO Signature of generic benchmark is rejected
/*
func BenchmarkFindString(b *testing.B, s string) {
	for n := 0; n < b.N; n++ {
		FindString(s)
	}
}
func BenchmarkFindStrings(b *testing.B) {
	for _, c := range testcases {
		BenchmarkFindString(b, c.search)
	}
}
*/
// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkFindStringShort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindString("bingo")
	}
}
func BenchmarkFindStringRangeShort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindStringRange("bingo")
	}
}
func BenchmarkFindStringLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindString("long")
	}
}
func BenchmarkFindStringRangeLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindStringRange("long")
	}
}
func BenchmarkFindStringNotFound(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindString("zorro")
	}
}
func BenchmarkFindStringRangeNotFound(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindStringRange("zorro")
	}
}
