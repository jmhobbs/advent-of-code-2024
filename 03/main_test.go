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
		assert.Equal(t, test.expected, main.ExtractMultipliers([]byte(test.input)))
	}
}

func Test_SumMultipliersFromCorruptedMemory(t *testing.T) {
	// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
	input := []byte("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	// ...produces 161 (2*4 + 5*5 + 11*8 + 8*5)
	assert.Equal(t, 161, main.SumMultipliersFromCorruptedMemory(input))
}
