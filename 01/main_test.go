package main_test

import (
	"strings"
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/01"

	"github.com/stretchr/testify/assert"
)

func Test_ListDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 11, main.ListDistance(left, right))
}

func Test_ListSimilarity(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 31, main.ListSimilarity(left, right))
}

func Test_ParseInput(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		input := `1   2
3   4
5   6`
		left, right, err := main.ParseInput(strings.NewReader(input))
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 3, 5}, left)
		assert.Equal(t, []int{2, 4, 6}, right)
	})

	t.Run("invalid left value", func(t *testing.T) {
		input := "notanumber   1"
		_, _, err := main.ParseInput(strings.NewReader(input))
		assert.Error(t, err)
	})

	t.Run("invalid right value", func(t *testing.T) {
		input := "1   parsethissucker"
		_, _, err := main.ParseInput(strings.NewReader(input))
		assert.Error(t, err)
	})
}
