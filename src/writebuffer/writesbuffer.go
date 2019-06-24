// writebuffer evaluates the difference between the elementary io.WriteString and through fmt.
package writebuffer

import (
	"bytes"
	"fmt"
	"log"
)

const s = "writestring"

func newBuffer() *bytes.Buffer {
	b := new(bytes.Buffer)
	_, _ = fmt.Fprint(b, s)
	return b
}

// fmt Package
func bufferFmt(ba *bytes.Buffer) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = fmt.Fprint(b, ba)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return b
}

// buffer Package
func bufferWriteTo(ba *bytes.Buffer) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = b.Write(ba.Bytes())
		if err != nil {
			log.Fatalln(err)
		}
	}
	return b
}
