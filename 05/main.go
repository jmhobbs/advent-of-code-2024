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

	rules, updates, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	middleSum, incorrectMiddleSum := ValidateOrSortUpdatesAndSumMiddles(rules, updates)

	fmt.Printf("A: %d\n", middleSum)
	fmt.Printf("B: %d\n", incorrectMiddleSum)
}

func ValidateOrSortUpdatesAndSumMiddles(rules []Rule, updates []Update) (int, int) {
	var (
		middleSum          int
		incorrectMiddleSum int
	)

	for _, update := range updates {
		if UpdateValid(rules, update) {
			middleSum += UpdateMiddle(update)
		} else {
			incorrectMiddleSum += UpdateMiddle(SortUpdate(rules, update))
		}
	}

	return middleSum, incorrectMiddleSum
}

type Rule struct {
	Before string
	After  string
}

type Update []string

func ParseInput(in io.Reader) ([]Rule, []Update, error) {
	rules := []Rule{}
	updates := []Update{}

	var (
		line  string
		split []string
	)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line = scanner.Text()

		split = strings.Split(line, "|")
		if len(split) == 2 {
			rules = append(rules, Rule{split[0], split[1]})
			continue
		}

		split = strings.Split(line, ",")
		if len(split) > 1 {
			updates = append(updates, Update(split))
		}
	}

	return rules, updates, scanner.Err()
}

func UpdateValid(rules []Rule, update Update) bool {
	for i, page := range update {
		for _, rule := range rules {
			if rule.After == page {
				if slices.Contains(update[i+1:], rule.Before) {
					return false
				}
			}
		}
	}
	return true
}

func UpdateMiddle(update Update) int {
	middle, err := strconv.ParseInt(update[len(update)/2], 10, 64)
	if err != nil {
		panic(err)
	}
	return int(middle)
}

func SortUpdate(rules []Rule, update Update) Update {
	slices.SortFunc(update, func(a, b string) int {
		for _, rule := range rules {
			if rule.After == a && rule.Before == b {
				return 1
			} else if rule.Before == a && rule.After == b {
				return -1
			}
		}
		return 0
	})
	return update
}
