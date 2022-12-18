package day17_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2022/day17"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

func SetUp() string {
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

func Test_Part1(t *testing.T) {
	line := SetUp()
	expected := 3068

	val, err := day.Part1Val(line)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	line := SetUp()
	expected := 0

	val, err := day.Part2Val(line)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
