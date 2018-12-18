package returnbuffer

import (
	"bytes"
	"fmt"
	"strconv"
)

const filling = "Filling the buffer "

func ReturnBuffer(f int) *bytes.Buffer {
	b := new(bytes.Buffer)
	for i := 1; i < f; i++ {
		if c, err := fmt.Fprint(b, filling, i); err != nil {
			fmt.Println("filling buffer failed")
		} else if c != len(filling)+len(strconv.Itoa(i)) {
			fmt.Printf("incomplete sentence written (*bytes.Buffer): got %d, want %d\n", c, len(filling))
		}
	}
	return b
}

func ReturnBufferString(f int) string {
	b := new(bytes.Buffer)
	for i := 1; i < f; i++ {
		if c, err := fmt.Fprint(b, filling, i); err != nil {
			fmt.Println("filling buffer failed")
		} else if c != len(filling)+len(strconv.Itoa(i)) {
			fmt.Printf("incomplete sentence written (.String): got %d, want %d\n", c, len(filling))
		}
	}
	return b.String()
}

func ReturnBufferBytes(f int) []byte {
	b := new(bytes.Buffer)
	for i := 1; i < f; i++ {
		if c, err := fmt.Fprint(b, filling, i); err != nil {
			fmt.Println("filling buffer failed")
		} else if c != len(filling)+len(strconv.Itoa(i)) {
			fmt.Printf("incomplete sentence written (.Bytes): got %d, want %d\n", c, len(filling))
		}
	}
	return b.Bytes()
}
