package util

import (
	"io"
	"os"
)

func OpenInput() io.ReadCloser {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	return f
}

func ReadInput() []byte {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return buf
}
