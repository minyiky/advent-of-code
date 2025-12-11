package day11_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2025/day11"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

//go:embed input_test_2.txt
var input2 string

func SetUp(input string) []string {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	return lines
}

func Test_Part1(t *testing.T) {
	lines := SetUp(input)
	expected := 5

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp(input2)
	expected := 2

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
