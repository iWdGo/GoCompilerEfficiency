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
	var err error
	for _, _ = range s {
		_, err = fmt.Fprint(b, s)
		if err != nil {
			log.Fatalln(err)
		}
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

// buffer package
func bufferWriteString(s string) *bytes.Buffer {
	b := new(bytes.Buffer)
	var err error
	for _, _ = range s {
		_, err = b.WriteString(s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return b
}

// TODO using defaunt print() after re-directing StdOut
