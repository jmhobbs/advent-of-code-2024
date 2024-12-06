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

type Rule struct {
	Before string
	After  string
}

type Update []string

func ParseInput(in io.Reader) ([]Rule, []Update, error) {
	rules := []Rule{}
	updates := []Update{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// nop
	}

	return rules, updates, scanner.Err()
}

func UpdateValid(rules []Rule, update Update) bool {
	return false
}

func UpdateMiddle(update Update) int {
	return 0
}
