package day08_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2023/day08"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

//go:embed input2_test.txt
var input2 string

//go:embed input3_test.txt
var input3 string

func SetUp(input string) []string {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	return lines
}

func Test_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{{
		input:    input,
		expected: 2,
	}, {
		input:    input2,
		expected: 6,
	}}
	for _, tc := range testCases {
		tc := tc
		t.Run("test", func(t *testing.T) {
			lines := SetUp(tc.input)

			val, err := day.Part1Val(lines)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, val)
		})
	}
}

func Test_Part2(t *testing.T) {
	lines := SetUp(input3)
	expected := 6

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
