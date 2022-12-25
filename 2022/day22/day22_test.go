package day22_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/minyiky/advent-of-code/2022/aocutils"
	day "github.com/minyiky/advent-of-code/2022/day22"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

func SetUp() []string {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	return lines
}

func GetFirst(line string) aocutils.Vector {
	var first aocutils.Vector
	for j, char := range []rune(line) {
		if char == '.' {
			first = aocutils.NewVector(j+1, 1)
			break
		}
	}
	return first
}

func ExtractGrid(lines []string) (map[aocutils.Vector]bool, map[aocutils.Vector]bool) {
	grid := make(map[aocutils.Vector]bool)
	blocks := make(map[aocutils.Vector]bool)

	for i, line := range lines {
		if line == "" {
			break
		}
		for j, char := range []rune(line) {
			switch char {
			case '#':
				blocks[aocutils.NewVector(j+1, i+1)] = true
				fallthrough
			case '.':
				grid[aocutils.NewVector(j+1, i+1)] = true
			}
		}
	}

	return grid, blocks
}

func Test_Part1(t *testing.T) {
	lines := SetUp()
	expected := 6032

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 5031

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
