package day07_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2022/day07"
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
	expected := 95437

	val := day.Part1Val(lines)

	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 24933642

	val, _ := day.Part2Val(lines)

	assert.Equal(t, expected, val)
}
