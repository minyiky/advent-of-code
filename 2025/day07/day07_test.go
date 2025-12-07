package day07_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/minyiky/advent-of-code/2025/day07"
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
	expected := 21

	val, err := day.Part1Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Test_Part2(t *testing.T) {
	lines := SetUp()
	expected := 40

	val, err := day.Part2Val(lines)

	assert.NoError(t, err)
	assert.Equal(t, expected, val)
}

func Benchmark_Part2_DFS(b *testing.B) {
	lines := SetUp()
	maxHeight, splitters, start := day.ParseInput(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day.SolveDFS(maxHeight, splitters, start)
	}
}

func Benchmark_Part2_BFS(b *testing.B) {
	lines := SetUp()
	maxHeight, splitters, start := day.ParseInput(lines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day.SolveBFS(maxHeight, splitters, start)
	}
}
