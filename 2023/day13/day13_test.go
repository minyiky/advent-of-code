package day13_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2023/day13"
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
	expected := 405

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 400

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
