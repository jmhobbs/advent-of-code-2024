package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	buf, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", SumMultipliersFromCorruptedMemory(buf))
}

type Mul struct {
	Left  int
	Right int
}

var regexpMul = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func ExtractMultipliers(input []byte) []Mul {
	muls := []Mul{}
	matches := regexpMul.FindAllSubmatch(input, -1)
	for _, match := range matches {
		left, err := strconv.ParseInt(string(match[1]), 10, 64)
		if err != nil {
			panic(err)
		}
		right, err := strconv.ParseInt(string(match[2]), 10, 64)
		if err != nil {
			panic(err)
		}
		muls = append(muls, Mul{int(left), int(right)})
	}
	return muls
}

func SumMultipliersFromCorruptedMemory(input []byte) int {
	sum := 0
	for _, m := range ExtractMultipliers(input) {
		sum += m.Left * m.Right
	}
	return sum
}

func SumMultipliersFromCorruptedMemoryWithEnablers(input []byte) int {
	return 0
}
