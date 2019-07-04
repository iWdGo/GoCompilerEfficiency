// writebuffer evaluates differences between fmt, io and buffer packages to write a buffer to another buffer
package writebuffer

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

const s = "writebuffer"

func newBuffer() *bytes.Buffer {
	// Simplest. Otherwise, you need to use another method like below.
	// No difference found.
	return bytes.NewBufferString(s)
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

// buffer Package requires to read Bytes() from the buffer.
func bufferWrite(ba *bytes.Buffer) *bytes.Buffer {
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

// buffer Package does not require conversion but to re-create the buffer.
func bufferWriteTo(ba *bytes.Buffer) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = ba.WriteTo(b)
		if err != nil {
			log.Fatalln(err)
		}
		// Because buffer to write ba is empty after WriteTo
		ba = bytes.NewBufferString(s)
	}
	return b
}

// buffer Package requires to read Bytes() from the buffer.
func ioCopyBuffer(ba *bytes.Buffer) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = io.CopyBuffer(b, ba, nil)
		if err != nil {
			log.Fatalln(err)
		}
		// Because buffer to write ba is empty after CopyBuffer
		ba = bytes.NewBufferString(s)
	}
	return b
}
