package main

import "testing"

var (
	data = []byte{3, 255}
	res  = []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
)

func arrayCompare(a1, a2 []int) bool {
	for i, a := range a1 {
		if a != a2[i] {
			return false
		}
	}
	return true
}

func TestBits(t *testing.T) {
	if d := forToBits(data); !arrayCompare(d, res) {
		t.Errorf("got %v, want %v", d, res)
	}
}

func TestRangeToBits(t *testing.T) {
	if d := rangeToBits(data); !arrayCompare(d, res) {
		t.Errorf("got %v, want %v", d, res)
	}
}

func TestRangeAndToBits(t *testing.T) {
	if d := andToBits(data); !arrayCompare(d, res) {
		t.Errorf("got %v, want %v", d, res)
	}
}

func TestRangeBitsToBits(t *testing.T) {
	if d := bitsToBits(data); !arrayCompare(d, res) {
		t.Errorf("got %v, want %v", d, res)
	}
}

func BenchmarkBits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		forToBits(data)
	}
}

func BenchmarkRangeToBits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rangeToBits(data)
	}
}

func BenchmarkAndToBits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		andToBits(data)
	}
}

func BenchmarkBitsToBits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bitsToBits(data)
	}
}
