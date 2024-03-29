package day15_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2023/day15"
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
	expected := 1320

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 145

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
