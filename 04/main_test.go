package main_test

import (
	"strings"
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/04"

	"github.com/stretchr/testify/assert"
)

func Test_ParseInput(t *testing.T) {
	actual, err := main.ParseInput(strings.NewReader(`..X...
.SAMX.
.A..A.
XMAS.S
.X....`))

	assert.Nil(t, err)
	assert.Equal(t, [][]byte{
		{'.', '.', 'X', '.', '.', '.'},
		{'.', 'S', 'A', 'M', 'X', '.'},
		{'.', 'A', '.', '.', 'A', '.'},
		{'X', 'M', 'A', 'S', '.', 'S'},
		{'.', 'X', '.', '.', '.', '.'},
	}, actual)
}

func Test_CountXmas(t *testing.T) {
	t.Run("forwards", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{{'.', 'X', 'M', 'A', 'S', '.'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'X', 'M', 'A', 'S', '.'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'.', 'X', 'M', 'A', 'S'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'X', 'M', 'A', 'S'}}))
		assert.Equal(t, 2, main.CountXmas([][]byte{{'X', 'M', 'A', 'S', 'X', 'M', 'A', 'S'}}))
	})

	t.Run("backwards", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{{'.', 'S', 'A', 'M', 'X', '.'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'S', 'A', 'M', 'X', '.'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'.', 'S', 'A', 'M', 'X'}}))
		assert.Equal(t, 1, main.CountXmas([][]byte{{'S', 'A', 'M', 'X'}}))
		assert.Equal(t, 2, main.CountXmas([][]byte{{'S', 'A', 'M', 'X', 'S', 'A', 'M', 'X'}}))
	})

	t.Run("up", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.'},
			{'.', 'X', '.'},
			{'.', 'M', '.'},
			{'.', 'A', '.'},
			{'.', 'S', '.'},
			{'.', '.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.'},
			{'.', 'X'},
			{'.', 'M'},
			{'.', 'A'},
			{'.', 'S'},
			{'.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.'},
			{'X', '.'},
			{'M', '.'},
			{'A', '.'},
			{'S', '.'},
			{'.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', 'X', '.'},
			{'.', 'M', '.'},
			{'.', 'A', '.'},
			{'.', 'S', '.'},
			{'.', '.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.'},
			{'.', 'X', '.'},
			{'.', 'M', '.'},
			{'.', 'A', '.'},
			{'.', 'S', '.'},
		}))
	})

	t.Run("down", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.'},
			{'.', 'S', '.'},
			{'.', 'A', '.'},
			{'.', 'M', '.'},
			{'.', 'X', '.'},
			{'.', '.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.'},
			{'.', 'S'},
			{'.', 'A'},
			{'.', 'M'},
			{'.', 'X'},
			{'.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.'},
			{'S', '.'},
			{'A', '.'},
			{'M', '.'},
			{'X', '.'},
			{'.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', 'S', '.'},
			{'.', 'A', '.'},
			{'.', 'M', '.'},
			{'.', 'X', '.'},
			{'.', '.', '.'},
		}))
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.'},
			{'.', 'S', '.'},
			{'.', 'A', '.'},
			{'.', 'M', '.'},
			{'.', 'X', '.'},
		}))
	})

	t.Run("up-left", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'S', '.', '.', '.'},
			{'.', 'A', '.', '.'},
			{'.', '.', 'M', '.'},
			{'.', '.', '.', 'X'},
		}))
	})

	t.Run("up-right", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.', 'S'},
			{'.', '.', 'A', '.'},
			{'.', 'M', '.', '.'},
			{'X', '.', '.', '.'},
		}))
	})

	t.Run("down-left", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'.', '.', '.', 'X'},
			{'.', '.', 'M', '.'},
			{'.', 'A', '.', '.'},
			{'S', '.', '.', '.'},
		}))
	})

	t.Run("down-right", func(t *testing.T) {
		assert.Equal(t, 1, main.CountXmas([][]byte{
			{'X', '.', '.', '.'},
			{'.', 'M', '.', '.'},
			{'.', '.', 'A', '.'},
			{'.', '.', '.', 'S'},
		}))
	})

	puzzle := [][]byte{
		{'.', '.', 'X', '.', '.', '.'},
		{'.', 'S', 'A', 'M', 'X', '.'},
		{'.', 'A', '.', '.', 'A', '.'},
		{'X', 'M', 'A', 'S', '.', 'S'},
		{'.', 'X', '.', '.', '.', '.'},
	}

	assert.Equal(t, 4, main.CountXmas(puzzle))

	puzzle = [][]byte{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	assert.Equal(t, 18, main.CountXmas(puzzle))
}

func Test_CountXMas(t *testing.T) {
	puzzle := [][]byte{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}

	assert.Equal(t, 1, main.CountXMas(puzzle))

	puzzle = [][]byte{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	assert.Equal(t, 9, main.CountXMas(puzzle))
}
