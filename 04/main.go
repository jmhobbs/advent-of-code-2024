package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = ParseInput(f)
	if err != nil {
		panic(err)
	}
}

func ParseInput(in io.Reader) ([][]byte, error) {
	lines := [][]byte{}

	var buf []byte
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		buf = scanner.Bytes()
		line := make([]byte, len(buf))
		copy(line, buf)
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func CountXmas(puzzle [][]byte) int {
	return 0
}
