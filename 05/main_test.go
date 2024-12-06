package main_test

import (
	"strings"
	"testing"

	main "github.com/jmhobbs/advent-of-code-2024/05"

	"github.com/stretchr/testify/assert"
)

func Test_ParseInput(t *testing.T) {
	rules, updates, err := main.ParseInput(strings.NewReader(`
123|456
456|789

123,456,789
123,789,456
`))
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]main.Rule{
			{"123", "456"},
			{"456", "789"},
		},
		rules,
	)
	assert.Equal(
		t,
		[]main.Update{
			{"123", "456", "789"},
			{"123", "789", "456"},
		},
		updates,
	)
}

func Test_UpdateValid(t *testing.T) {
	rules := []main.Rule{
		{"47", "53"},
		{"97", "13"},
		{"97", "61"},
		{"97", "47"},
		{"75", "29"},
		{"61", "13"},
		{"75", "53"},
		{"29", "13"},
		{"97", "29"},
		{"53", "29"},
		{"61", "53"},
		{"97", "53"},
		{"61", "29"},
		{"47", "13"},
		{"75", "47"},
		{"97", "75"},
		{"47", "61"},
		{"75", "61"},
		{"47", "29"},
		{"75", "13"},
		{"53", "13"},
	}

	// In the above example, the first update (`75,47,61,53,29`) is in the right order:
	t.Run("75,47,61,53,29", func(t *testing.T) {
		assert.True(t, main.UpdateValid(rules, main.Update{"75", "47", "61", "53", "29"}))
	})

	// The second and third updates are also in the correct order according to the rules.
	t.Run("97,61,53,29,13", func(t *testing.T) {
		assert.True(t, main.UpdateValid(rules, main.Update{"97", "61", "53", "29", "13"}))
	})
	t.Run("75,29,13", func(t *testing.T) {
		assert.True(t, main.UpdateValid(rules, main.Update{"75", "29", "13"}))
	})

	// The fourth update, `75,97,47,61,53`, is not in the correct order:
	t.Run("75,97,47,61,53", func(t *testing.T) {
		assert.False(t, main.UpdateValid(rules, main.Update{"75", "97", "47", "61", "53"}))
	})

	// The fifth update, 61,13,29, is also not in the correct order
	t.Run("61,13,29", func(t *testing.T) {
		assert.False(t, main.UpdateValid(rules, main.Update{"61", "13", "29"}))
	})

	// The last update, `97,13,75,29,47`, is not in the correct order due to breaking several rules.
	t.Run("97,13,75,29,47", func(t *testing.T) {
		assert.False(t, main.UpdateValid(rules, main.Update{"97", "13", "75", "29", "47"}))
	})
}

func Test_UpdateMiddle(t *testing.T) {
	// 75,47,61,53,29
	// 97,61,53,29,13
	// 75,29,13
	// These have middle page numbers of 61, 53, and 29 respectively. Adding these page numbers together gives 143.
	assert.Equal(t, 61, main.UpdateMiddle(main.Update{"75", "47", "61", "53", "29"}))
	assert.Equal(t, 53, main.UpdateMiddle(main.Update{"97", "61", "53", "29", "13"}))
	assert.Equal(t, 29, main.UpdateMiddle(main.Update{"75", "29", "13"}))
}
