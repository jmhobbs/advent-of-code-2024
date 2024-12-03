package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	assert.Equal(t, 5, Abs(5))
	assert.Equal(t, 5, Abs(-5))
	assert.Equal(t, 0, Abs(0))
}
