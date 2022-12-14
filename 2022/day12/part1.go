package day12

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func checkHeight(current, next rune) bool {
	return next <= (current + 1)
}

func Part1Val(lines []string) (int, error) {
	yLen, xLen := len(lines), len(lines[1])

	grid := make([][]rune, yLen)

	for i := range grid {
		grid[i] = make([]rune, xLen)
	}

	var start, end aocutils.Vector
	for y, line := range lines {
		for x, char := range []rune(line) {
			if char == 'S' {
				start = aocutils.NewVector(x, y)
				grid[y][x] = 'a'
				continue
			}
			if char == 'E' {
				end = aocutils.NewVector(x, y)
				grid[y][x] = 'z'
				continue
			}
			grid[y][x] = char
		}
	}

	emptyMap := make(map[aocutils.Vector]int)

	value, _ := findSummit(start, end, 0, grid, emptyMap)

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Using the shortest route it would take  %d steps to reach the summit\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
