package day02_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/minyiky/advent-of-code/2022/day02"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

func SetUp() []string {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	return lines
}

func Test_Part1(t *testing.T) {
	lines := SetUp()
	expected := 15

	val, err := day02.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 12

	val, err := day02.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
