package forloops

import (
	"log"
	"math"
	"testing"
)

var tests = []struct {
	i    int
	want bool
}{
	{2, true},
	{3, false},
	{1024, true}, // math.Pow(2,10)
	{1024 + 1, false},
	{int(math.Pow(2, 31)), true}, // max int
}

func TestIs2SingleDivBig(t *testing.T) {
	for _, test := range tests {
		if got := is2SingleDivBig(test.i); got != test.want {
			t.Fatalf("Is2SingleDivBig(%d): want %v, got %v\n", test.i, test.want, got)
		}
	}

}
func TestIs2SingleDivLog(t *testing.T) {
	for _, test := range tests {
		if got := is2SingleDivLog(test.i); got != test.want {
			t.Fatalf("Is2SingleDivLog(%d):want %v, got %v\n", test.i, test.want, got)
		}
	}

}
func TestIs2SingleDivDivision(t *testing.T) {
	for _, test := range tests {
		if got := is2SingleDivDivision(test.i); got != test.want {
			t.Fatalf("Is2SingleDivDivision(%d): want %v, got %v\n", test.i, test.want, got)
		}
	}
}
func TestIs2SingleDivShift(t *testing.T) {
	for _, test := range tests {
		if got := is2SingleDivShift(test.i); got != test.want {
			t.Fatalf("Is2SingleDivDivision(%d): want %v, got %v\n", test.i, test.want, got)
		}
	}
}

// Any divisor
func TestIsSingleDivLog(t *testing.T) {
	for _, test := range tests {
		if test.i < 1073741824 {
			if got := isSingleDivLog(test.i, 2); got != test.want {
				t.Fatalf("Is2SingleDivLog(%d):want %v, got %v\n", test.i, test.want, got)
			}
		}
	}

}
func TestIsSingleDivDivision(t *testing.T) {
	for _, test := range tests {
		if got := isSingleDivDivision(test.i, 2); got != test.want {
			t.Fatalf("Is2SingleDivDivision(%d): want %v, got %v\n", test.i, test.want, got)
		}
	}
}

// Benchmark
// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkIs2SingleDivBig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !is2SingleDivBig(1024) {
			log.Fatal("Benchmark failed")
		}
	}
}
func BenchmarkIs2SingleDivLog(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !is2SingleDivLog(1024) {
			log.Fatal("Benchmark failed")
		}
	}
}
func BenchmarkIs2SingleDivDivision(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !is2SingleDivDivision(1024) {
			log.Fatal("Benchmark failed")
		}
	}
}
func BenchmarkIs2SingleDivShift(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !is2SingleDivShift(1024) {
			log.Fatal("Benchmark failed")
		}
	}
}

//1073741824
func BenchmarkIsSingleDivLog(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !isSingleDivLog(1024, 2) {
			log.Fatal("Benchmark failed")
		}
	}
}
func BenchmarkIsSingleDivDivision(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if !isSingleDivDivision(1024, 2) {
			log.Fatal("Benchmark failed")
		}
	}
}
