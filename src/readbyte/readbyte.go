package readbyte

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

const (
	filling = "Filling the buffer "
	times   = 20
	want    = len(filling)*times + 9 + 2*(times-9) // 9 one digit + 11 2-digits (11 to 19)
)

func fillBuffer(f int) *bytes.Buffer {
	b := new(bytes.Buffer)
	for i := 1; i <= f; i++ {
		if c, err := fmt.Fprintf(b, "%s%d", filling, i); err != nil {
			log.Fatalf("filling buffer failed")
		} else if c != len(filling)+len(strconv.Itoa(i)) {
			log.Fatalf("incomplete sentence written (*bytes.Buffer): got %d, want %d\n", c, len(filling))
		}
	}
	if b.Len() != want {
		log.Fatalf("filling buffer failed: got %d, want %d\n", b.Len(), want)
	}
	return b
}

func ReadAllBytesRead(b *bytes.Buffer) (l int) {
	var err error
	var n int // avoid shadowing
	gotb := make([]byte, 1)
	for err != io.EOF {
		n, err = b.Read(gotb)
		l += n
	}
	return
}

func ReadAllBytesReadByte(b *bytes.Buffer) (l int) {
	var err error
	var gotb, e byte
	for err != io.EOF {
		gotb, err = b.ReadByte()
		if gotb != e {
			l++
		}
	}
	return
}
