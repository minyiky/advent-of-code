package day06_test

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2022/day06"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

func SetUp() []string {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n\n")
	return lines
}

func Test_Part1(t *testing.T) {
	lines := SetUp()
	lines = strings.Split(lines[0], "\n")
	for i, line := range lines {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			parts := strings.Split(line, " ")
			expected, _ := strconv.Atoi(parts[1])

			val, err := day.Part1Val(parts[0])

			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		})
	}
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	lines = strings.Split(lines[1], "\n")
	for i, line := range lines {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			parts := strings.Split(line, " ")
			expected, _ := strconv.Atoi(parts[1])

			val, err := day.Part2Val(parts[0])

			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		})
	}
}
