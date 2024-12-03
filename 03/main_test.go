package main_test

import (
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/03"

	"github.com/stretchr/testify/assert"
)

func Test_ExtractOperations(t *testing.T) {
	tests := []struct {
		input    string
		expected []main.Op
	}{
		// `mul(44,46)` multiplies `44` by `46` to get a result of `2024`.
		{
			"mul(44,46)",
			[]main.Op{{main.Mul, 44, 46}},
		},
		// Similarly, `mul(123,4)` would multiply `123` by `4`.
		{
			"mul(123,4)",
			[]main.Op{{main.Mul, 123, 4}},
		},
		// `mul(4*`, `mul(6,9!`, `?(12,34)`, or `mul ( 2 , 4 )` do nothing.
		{
			"mul(4*",
			[]main.Op{},
		},
		{
			"mul(6,9!",
			[]main.Op{},
		},
		{
			"?(12,34)",
			[]main.Op{},
		},
		{
			"mul ( 2 , 4 )",
			[]main.Op{},
		},
		{
			// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			// ...produces 161 (2*4 + 5*5 + 11*8 + 8*5)
			[]main.Op{
				{main.Mul, 2, 4},
				{main.Mul, 5, 5},
				{main.Mul, 11, 8},
				{main.Mul, 8, 5},
			},
		},
		{
			// For example:
			//   xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			// This time, the sum of the results is 48 (2*4 + 8*5).
			[]main.Op{
				{main.Mul, 2, 4},
				{main.Dont, 0, 0},
				{main.Mul, 5, 5},
				{main.Mul, 11, 8},
				{main.Do, 0, 0},
				{main.Mul, 8, 5},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, main.ExtractOperations([]byte(test.input)))
	}
}

func Test_SumMultipliers(t *testing.T) {
	// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
	input :=
		[]main.Op{
			{main.Mul, 2, 4},
			{main.Mul, 5, 5},
			{main.Mul, 11, 8},
			{main.Mul, 8, 5},
		}
	// ...produces 161 (2*4 + 5*5 + 11*8 + 8*5)
	assert.Equal(t, 161, main.SumMultipliers(input))
}

func Test_SumMultipliersWithEnablers(t *testing.T) {
	// For example:
	//   xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
	input := []main.Op{
		{main.Mul, 2, 4},
		{main.Dont, 0, 0},
		{main.Mul, 5, 5},
		{main.Mul, 32, 64},
		{main.Mul, 11, 8},
		{main.Do, 0, 0},
		{main.Mul, 8, 5},
	}
	// This time, the sum of the results is 48 (2*4 + 8*5).
	assert.Equal(t, 48, main.SumMultipliersWithEnablers(input))
}
