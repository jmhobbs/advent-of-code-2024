package main

import (
	"bufio"
	"fmt"
	"os"
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
		safeReports             int
		safeReportsWithDampener int
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levelStrs := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(levelStrs))
		for i, v := range levelStrs {
			levels[i], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}
		if InputsSafe(levels) {
			safeReports += 1
		}
		if InputsSafeWithDampener(levels) {
			safeReportsWithDampener += 1
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", safeReports)
	fmt.Printf("B: %d\n", safeReportsWithDampener)
}

func InputsSafeWithDampener(levels []int) bool {
	if InputsSafe(levels) {
		return true
	}

	for i := range levels {
		if InputsSafe(append(append([]int{}, levels[:i]...), levels[i+1:]...)) {
			return true
		}
	}

	return false
}

func InputsSafe(levels []int) bool {
	var ascending bool

	for i, v := range levels {
		if i == 0 {
			ascending = levels[1] > v
			continue
		}

		if ascending {
			if levels[i-1] >= v {
				return false
			}
		} else {
			if levels[i-1] <= v {
				return false
			}
		}

		if abs(v-levels[i-1]) > 3 {
			return false
		}
	}

	return true
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
