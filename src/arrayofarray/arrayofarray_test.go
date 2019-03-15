package arrayofarray

import (
	"testing"
)

const (
	inc  = 10
	want = 10 * size
)

// [][]int
func llen(ab [][]int) (l int) {
	for _, a := range ab {
		l += len(a)
	}
	return
}
func TestInitArray(t *testing.T) {
	want := want - size*(size-1)/2
	if got := initArray(); llen(got) != want {
		t.Fatalf("want %d, got %d\n %v", want, llen(got), got)
	}
}
func TestUpdateArray(t *testing.T) {
	want := want - size*(size-1)/2 + inc*size
	if got := updateArray(initArray(), inc); llen(got) != want {
		t.Fatalf("want %d, got %d\n %v", want, llen(got), got)
	} else {
		for _, a := range got {
			if gotv := a[len(a)-inc]; gotv != 0 {
				t.Fatalf("want 0, got %d in %v", gotv, a)
			}
		}
	}

}

// []t
func llent(ab []t) (l int) {
	for _, a := range ab {
		l += len(a)
	}

	return
}
func TestUpdateArrayType(t *testing.T) {
	want := want - size*(size-1)/2 + inc*size
	if got := updateArrayType(initArrayType(), inc); llent(got) != want {
		t.Fatalf("want %d, got %d\n %v", want, llent(got), got)
	} else {
		for _, a := range got {
			if gotv := a[len(a)-inc]; gotv != 0 {
				t.Fatalf("want 0, got %d in %v", gotv, a)
			}
		}
	}
}

// tt
func (ab tt) llen() (l int) {
	for _, a := range ab {
		l += len(a)
	}
	return
}
func TestInitArrayType(t *testing.T) {
	want := want - size*(size-1)/2
	if got := initArrayTyped(); got.llen() != want {
		t.Fatalf("want %d, got %d\n %v", want, got.llen(), got)
	}
}
func TestUpdateArrayMethod(t *testing.T) {
	var tb tt
	tb = initArrayType()
	want := want - size*(size-1)/2 + inc*size
	tb.updateArrayMethod(inc)
	if ll := tb.llen(); ll != want {
		t.Fatalf("want %d, got %d\n %v", want, ll, tb)
	} else {
		for _, a := range tb {
			if gotv := a[len(a)-inc]; gotv != 0 {
				t.Fatalf("want 0, got %d in %v", gotv, a)
			}
		}
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkUpdateArray(b *testing.B) {
	tb := initArray()
	for n := 0; n < b.N; n++ {
		updateArray(tb, inc)
	}
}
func BenchmarkUpdateArrayType(b *testing.B) {
	tb := initArrayType()
	for n := 0; n < b.N; n++ {
		updateArrayType(tb, inc)
	}
}
func BenchmarkUpdateArrayMethod(b *testing.B) {
	var tb tt
	tb = initArrayType()
	for n := 0; n < b.N; n++ {
		tb.updateArrayMethod(inc)
	}
}
