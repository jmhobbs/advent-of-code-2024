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

	_, _, err = ParseInput(f)
	if err != nil {
		panic(err)
	}
}

type Robot struct {
	X, Y        int
	Orientation byte
}

type Map [][]byte

func ParseInput(in io.Reader) (Map, Robot, error) {
	lines := [][]byte{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}

	return Map(lines), Robot{0, 0, '?'}, scanner.Err()
}

func Step(robot Robot, mmap Map) bool {
	return false
}

func WalkMap(robot Robot, mmap Map) {}

func CountVisitedPositions(mmap Map) int {
	return 0
}
