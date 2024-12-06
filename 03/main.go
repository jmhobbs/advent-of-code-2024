package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/jmhobbs/advent-of-code-2024/util"
)

func main() {
	ops := ExtractOperations(util.ReadInput())

	fmt.Printf("A: %d\n", SumMultipliers(ops))
	fmt.Printf("B: %d\n", SumMultipliersWithEnablers(ops))
}

type OpName uint8

const (
	Do OpName = iota
	Dont
	Mul
)

type Op struct {
	Name  OpName
	Left  int
	Right int
}

var regexpMul = regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)

func ExtractOperations(input []byte) []Op {
	ops := []Op{}
	matches := regexpMul.FindAllSubmatch(input, -1)
	for _, match := range matches {
		if bytes.Equal(match[1], []byte("do()")) {
			ops = append(ops, Op{Do, 0, 0})
			continue
		}
		if bytes.Equal(match[1], []byte("don't()")) {
			ops = append(ops, Op{Dont, 0, 0})
			continue
		}
		left, err := strconv.ParseInt(string(match[2]), 10, 64)
		if err != nil {
			panic(err)
		}
		right, err := strconv.ParseInt(string(match[3]), 10, 64)
		if err != nil {
			panic(err)
		}
		ops = append(ops, Op{Mul, int(left), int(right)})
	}
	return ops
}

func SumMultipliers(input []Op) int {
	sum := 0
	for _, m := range input {
		if m.Name == Mul {
			sum += m.Left * m.Right
		}
	}
	return sum
}

func SumMultipliersWithEnablers(input []Op) int {
	sum := 0
	enabled := true
	for _, m := range input {
		switch m.Name {
		case Do:
			enabled = true
		case Dont:
			enabled = false
		case Mul:
			if enabled {
				sum += m.Left * m.Right
			}
		}
	}
	return sum
}
