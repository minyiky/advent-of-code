package day15_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
	day "github.com/minyiky/advent-of-code/2022/day15"
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
	expected := 26
	row := 10

	val, err := day.Part1Val(lines, row)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 56000011

	bottomLeft := point.NewPoint2D(0, 0)
	topRight := point.NewPoint2D(20, 20)

	val, err := day.Part2Val(lines, &bottomLeft, &topRight)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
