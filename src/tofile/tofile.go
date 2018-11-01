package tmplfile

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	design = `This content is read from a const
and treated like a string. From there, it is moved in the
benmarked structure:
- string
- bytes.Buffer
- io.ReadCloser

The benchmark is about the speed of transfer to the 
file system not the file system. So the file is re-created (content.txt)
and is written to in each test.
`

	filename = "content.txt"
)

func StringToFile(fname string, content string) (n int) {
	wfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	// defer wfile.Close()
	if err != nil {
		panic(err)
	}

	n, err = wfile.WriteString(content) // ? sometimes the opening < of body is not saved ?
	if err != nil {
		panic(err)
	}
	wfile.Close() // not defered but executed when processing is complete
	return
}

func BufferToFile(fname string, content *bytes.Buffer) (n int) {
	wfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	// defer wfile.Close()
	if err != nil {
		panic(err)
	}

	n, err = wfile.Write(content.Bytes()) // ? sometimes the opening < of body is not saved ?
	if err != nil {
		panic(err)
	}
	wfile.Close()
	return
}
func ReadCloserToFileOneRead(fname string, content io.ReadCloser) (n int) {
	var err error
	wfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	c, err := ioutil.ReadAll(content)
	if err != nil {
		panic(err)
	}
	n, err = wfile.Write(c)
	if err != nil {
		panic(err)
	}
	if len(c) != n {
		fmt.Printf("file %s is missing %d bytes", fname, len(c)-n)
	}
	wfile.Close()
	return
}

func ReadCloserToFileLoop(fname string, content io.ReadCloser) (n int) {
	var err error
	wfile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	c := 0
	b := make([]byte, 1)
	for err != nil {
		c, err = content.Read(b)
		if err != nil && err != io.EOF {
			panic(err)
		} else if c == 1 {
			fmt.Printf("%s", b)
			_, err = wfile.Write(b)
			if err != nil {
				panic(err)
			}
		}
		n++
	}
	wfile.Close()
	return
}
