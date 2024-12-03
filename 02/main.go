package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/jmhobbs/advent-of-code-2024/util"
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

	levelInputs, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	for _, levels := range levelInputs {
		if InputsSafe(levels) {
			safeReports += 1
		}
		if InputsSafeWithDampener(levels) {
			safeReportsWithDampener += 1
		}
	}

	fmt.Printf("A: %d\n", safeReports)
	fmt.Printf("B: %d\n", safeReportsWithDampener)
}

func ParseInput(input io.Reader) ([][]int, error) {
	var (
		levelInputs [][]int = [][]int{}
		err         error
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		levelStrs := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(levelStrs))
		for i, v := range levelStrs {
			levels[i], err = strconv.Atoi(v)
			if err != nil {
				return levelInputs, err
			}
		}
		levelInputs = append(levelInputs, levels)
	}
	return levelInputs, scanner.Err()
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

		if util.Abs(v-levels[i-1]) > 3 {
			return false
		}
	}

	return true
}
