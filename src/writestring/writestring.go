// writestring evaluates the difference between the elementary io.WriteString and through fmt.
package writestring

import (
	"bytes"
	"fmt"
	"io"
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
	for _, _ = range s {
		_, _ = io.WriteString(b, s)
	}
	return b
}

// TODO using defaunt print() after re-directing StdOut
