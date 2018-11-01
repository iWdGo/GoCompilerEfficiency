package singleexit

import (
	"strings"
	"testing"
)

var testcases = []struct {
	want   bool
	search string
}{
	{true, "alpha"},
	{true, "bingo"},
	{true, "cause"},
	{true, "daemon"},
	{true, "enola"},
	{true, "finder"},
	{true, "goal"},
	{true, "helio"},
	{true, "iota"},
	{true, "jolien"},
	{true, "kilo"},
	{true, "long"},
	{true, "mono"},
	{true, "nano"}, // last found
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
func BenchmarkMain(b *testing.B) {
	for _, c := range testcases {
		b.Run(c.search+"/while", BenchmarkFind)
		b.Run(c.search+"/range", BenchmarkFind)
	}
}
func BenchmarkFind(b *testing.B) {
	s := strings.SplitAfter(b.Name(), "/")
	// because BenchMarkFind is called w/o suffix with >go test -bench=.
	if len(s) == 3 {
		if s[2] == "while" {
			for n := 0; n < b.N; n++ {
				FindString(s[1])
			}
		} else {
			for n := 0; n < b.N; n++ {
				FindStringRange(s[1])
			}
		}
	} else {
		// Benchmark name has no suffix
	}
}

/*
func BenchmarkFindStrings(b *testing.B) {
	for _, c := range testcases {
		BenchmarkFindString(b, c.search) // signature is rejected
	}
}
*/
