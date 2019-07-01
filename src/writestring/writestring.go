// writestring benchmarks the elementary io.WriteString and package fmt.
package writestring

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

const s = "writestring"

// fmt Package
func fmtWriteString(s string) *bytes.Buffer {
	b := new(bytes.Buffer)
	for _, _ = range s {
		_, _ = fmt.Fprint(b, s)
	}
	return b
}

// io Package
func ioWriteString(s string) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = io.WriteString(b, s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return b
}

// bytes Package
func bytesWriteTo(s string) *bytes.Buffer {
	b := new(bytes.Buffer)
	bs := bytes.NewBufferString(s)
	var err error
	for _, _ = range s {
		_, err = bs.WriteTo(b)
		if err != nil {
			log.Fatalln(err)
		}
		// ! the new buffer is (probably) re-used and has to be re-created
		// Several issues still opened for bytes.NewBufferString()
		// https://github.com/golang/go/issues?utf8=%E2%9C%93&q=is%3Aissue+newbufferstring
		bs = bytes.NewBufferString(s)
	}
	return b
}

// TODO using defaunt print() after re-directing StdOut
