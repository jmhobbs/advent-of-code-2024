package main

import (
	"bufio"
	"fmt"
	"io"
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

	rules, updates, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	var middleSum int
	for _, update := range updates {
		if UpdateValid(rules, update) {
			middleSum += UpdateMiddle(update)
		}
	}

	fmt.Printf("A: %d\n", middleSum)
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
