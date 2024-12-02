package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListDistances(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 11, ListDistances(left, right))
}
