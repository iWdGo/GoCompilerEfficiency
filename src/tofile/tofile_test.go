package tmplfile

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

/* */
func TestMain(m *testing.M) {
	// os.Remove(filename)
	if f, err := os.Create(filename); err != nil {
		log.Fatalln("can't create file", filename)
	} else {
		f.Close()
	}
	m.Run()
}
func TestStringToFile(t *testing.T) {
	if got := StringToFile(filename, design); len(design) != got {
		t.Fatalf("stringtofile: got %d, want %d bytes", got, len(design))
	}
}

func TestBufferToFile(t *testing.T) {
	bb := bytes.NewBufferString(design)
	if got := BufferToFile(filename, bb); len(design) != got {
		t.Fatalf("buffertofile: got %d, want %d bytes", got, len(design))
	}
}

func TestReadCloserToFile(t *testing.T) {
	i := ioutil.NopCloser(bytes.NewReader([]byte(design))) // io.ReadCloser containing your string
	if got := ReadCloserToFileOneRead(filename, i); len(design) != got {
		t.Fatalf("readcloser: got %d, want %d bytes", got, len(design))
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkStringToFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if got := StringToFile(filename, design); len(design) != got {
			b.Fatalf("stringtofile: got %d, want %d bytes", got, len(design))
		}
	}
}
func BenchmarkBufferToFile(b *testing.B) {
	bb := bytes.NewBufferString(design)
	for n := 0; n < b.N; n++ {
		if got := BufferToFile(filename, bb); len(design) != got {
			b.Fatalf("buffertofile: got %d, want %d bytes", got, len(design))
		}
	}
}

// TODO Benchmark fails
func BenchmarkReadCloserToFile(b *testing.B) {
	b.SkipNow()
	i := ioutil.NopCloser(bytes.NewReader([]byte(design))) // io.ReadCloser containing your string
	for n := 0; n < b.N; n++ {
		if got := ReadCloserToFileOneRead(filename, i); len(design) != got {
			b.Fatalf("readcloser: got %d, want %d bytes", got, len(design))
		}
	}
}
