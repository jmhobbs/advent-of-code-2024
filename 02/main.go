package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jmhobbs/advent-of-code-2024/util"
)

func main() {
	f := util.OpenInput()
	defer f.Close()

	reports, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("A: %d\n", CountSafeReports(reports))
	fmt.Printf("B: %d\n", CountSafeReportsWithDampener(reports))
}

func ParseInput(input io.Reader) ([][]int, error) {
	var (
		reports [][]int = [][]int{}
		err     error
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		levelStrs := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(levelStrs))
		for i, v := range levelStrs {
			levels[i], err = strconv.Atoi(v)
			if err != nil {
				return reports, err
			}
		}
		reports = append(reports, levels)
	}
	return reports, scanner.Err()
}

func CountSafeReports(reports [][]int) int {
	var safeReports int
	for _, levels := range reports {
		if InputsSafe(levels) {
			safeReports += 1
		}
	}
	return safeReports
}

func CountSafeReportsWithDampener(reports [][]int) int {
	var safeReports int
	for _, levels := range reports {
		if InputsSafeWithDampener(levels) {
			safeReports += 1
		}
	}
	return safeReports
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
