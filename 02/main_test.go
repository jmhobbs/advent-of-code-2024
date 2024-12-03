package main_test

import (
	"strings"
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

func Test_InputsSafeWithDampener(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		// `7 6 4 2 1`: Safe without removing any level.
		{
			[]int{7, 6, 4, 2, 1},
			true,
		},
		// `1 2 7 8 9`: Unsafe regardless of which level is removed.
		{
			[]int{1, 2, 7, 8, 9},
			false,
		},
		// `9 7 6 2 1`: Unsafe regardless of which level is removed.
		{
			[]int{9, 7, 6, 2, 1},
			false,
		},
		// `1 3 2 4 5`: Safe by removing the second level, 3.
		{
			[]int{1, 3, 2, 4, 5},
			true,
		},
		// `8 6 4 4 1`: Safe by removing the third level, 4.
		{
			[]int{8, 6, 4, 4, 1},
			true,
		},
		// `1 3 6 7 9`: Safe without removing any level.
		{
			[]int{1, 3, 6, 7, 9},
			true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, main.InputsSafeWithDampener(test.input))
	}
}

func Test_ParseInput(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
`
		levels, err := main.ParseInput(strings.NewReader(input))
		assert.NoError(t, err)
		assert.Equal(t, [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
		}, levels)
	})

	t.Run("invalid numbers", func(t *testing.T) {
		input := `7 6 4 2 1
1 no 7 8 9
9 7 6 2 1
`
		_, err := main.ParseInput(strings.NewReader(input))
		assert.Error(t, err)
	})

	t.Run("invalid spacing", func(t *testing.T) {
		input := `7 6  4 2 1`
		_, err := main.ParseInput(strings.NewReader(input))
		assert.Error(t, err)
	})
}
