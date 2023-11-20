package day22_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
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

func GetFirst(line string) point.Point2D {
	var first point.Point2D
	for j, char := range []rune(line) {
		if char == '.' {
			first = point.NewPoint2D(j+1, 1)
			break
		}
	}
	return first
}

func ExtractGrid(lines []string) (map[point.Point2D]bool, map[point.Point2D]bool) {
	grid := make(map[point.Point2D]bool)
	blocks := make(map[point.Point2D]bool)

	for i, line := range lines {
		if line == "" {
			break
		}
		for j, char := range []rune(line) {
			switch char {
			case '#':
				blocks[point.NewPoint2D(j+1, i+1)] = true
				fallthrough
			case '.':
				grid[point.NewPoint2D(j+1, i+1)] = true
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
