package day12

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
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

	var start, end point.Point2D
	for y, line := range lines {
		for x, char := range []rune(line) {
			if char == 'S' {
				start = point.NewPoint2D(x, y)
				grid[y][x] = 'a'
				continue
			}
			if char == 'E' {
				end = point.NewPoint2D(x, y)
				grid[y][x] = 'z'
				continue
			}
			grid[y][x] = char
		}
	}

	emptyMap := make(map[point.Point2D]int)

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
