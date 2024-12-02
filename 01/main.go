package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		left  []int
		right []int
		split []string
		value int64
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split = strings.SplitN(scanner.Text(), "   ", 2)
		value, err = strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		left = append(left, int(value))

		value, err = strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			panic(err)
		}
		right = append(right, int(value))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", ListDistance(left, right))
	fmt.Printf("B: %d\n", ListSimilarity(left, right))
}

func ListDistance(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	var distance int

	for i, l := range left {
		distance += abs(l - right[i])
	}

	return distance
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
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
