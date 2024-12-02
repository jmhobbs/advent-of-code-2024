package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 11, ListDistance(left, right))
}

func Test_ListSimilarity(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 31, ListSimilarity(left, right))
}
