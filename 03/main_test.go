package main_test

import (
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/03"

	"github.com/stretchr/testify/assert"
)

func Test_ExtractMultipliers(t *testing.T) {
	tests := []struct {
		input    string
		expected []main.Mul
	}{
		// `mul(44,46)` multiplies `44` by `46` to get a result of `2024`.
		{
			"mul(44,46)",
			[]main.Mul{{44, 46}},
		},
		// Similarly, `mul(123,4)` would multiply `123` by `4`.
		{
			"mul(123,4)",
			[]main.Mul{{123, 4}},
		},
		// `mul(4*`, `mul(6,9!`, `?(12,34)`, or `mul ( 2 , 4 )` do nothing.
		{
			"mul(4*",
			[]main.Mul{},
		},
		{
			"mul(6,9!",
			[]main.Mul{},
		},
		{
			"?(12,34)",
			[]main.Mul{},
		},
		{
			"mul ( 2 , 4 )",
			[]main.Mul{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, main.ExtractMultipliers(test.input))
	}
}
