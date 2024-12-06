package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/jmhobbs/advent-of-code-2024/util"
)

func main() {
	f := util.OpenInput()
	defer f.Close()

	left, right, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", ListDistance(left, right))
	fmt.Printf("B: %d\n", ListSimilarity(left, right))
}

func ParseInput(input io.Reader) ([]int, []int, error) {
	var (
		left  []int = []int{}
		right []int = []int{}
		split []string
		value int64
		err   error
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		split = strings.SplitN(scanner.Text(), "   ", 2)

		value, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return left, right, err
		}
		left = append(left, int(value))

		value, err = strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return left, right, err
		}
		right = append(right, int(value))
	}

	return left, right, scanner.Err()
}

func ListDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	var distance int

	for i, l := range left {
		distance += util.Abs(l - right[i])
	}

	return distance
}

func ListSimilarity(left []int, right []int) int {
	var (
		score   int
		counter int
	)

	for _, l := range left {
		counter = 0
		for _, r := range right {
			if l == r {
				counter++
			}
		}
		score += counter * l
	}

	return score
}
