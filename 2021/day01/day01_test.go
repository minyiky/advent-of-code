package day01_test

import (
	_ "embed"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2021/day01"
	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input string

func SetUp() []int {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	ints := make([]int, len(lines))
	for i, line := range lines {
		var err error
		ints[i], err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	return ints
}

func Test_Part1(t *testing.T) {
	lines := SetUp()
	expected := 7

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 5

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}
