package main_test

import (
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/02"

	"github.com/stretchr/testify/assert"
)

func Test_InputsSafe(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		// `7 6 4 2 1`: Safe because the levels are all decreasing by 1 or 2.
		{
			[]int{7, 6, 4, 2, 1},
			true,
		},
		// `1 2 7 8 9`: Unsafe because 2 7 is an increase of 5.
		{
			[]int{1, 2, 7, 8, 9},
			false,
		},
		// `9 7 6 2 1`: Unsafe because 6 2 is a decrease of 4.
		{
			[]int{9, 7, 6, 2, 1},
			false,
		},
		// `1 3 2 4 5`: Unsafe because 1 3 is increasing but 3 2 is decreasing.
		{
			[]int{1, 3, 2, 4, 5},
			false,
		},
		// `8 6 4 4 1`: Unsafe because 4 4 is neither an increase or a decrease.
		{
			[]int{8, 6, 4, 4, 1},
			false,
		},
		// `1 3 6 7 9`: Safe because the levels are all increasing by 1, 2, or 3.
		{
			[]int{1, 3, 6, 7, 9},
			true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, main.InputsSafe(test.input))
	}
}
