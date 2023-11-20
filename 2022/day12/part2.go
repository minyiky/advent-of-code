package day12

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	yLen, xLen := len(lines), len(lines[1])

	grid := make([][]rune, yLen)

	for i := range grid {
		grid[i] = make([]rune, xLen)
	}

	var start []point.Point2D
	var end point.Point2D
	for y, line := range lines {
		for x, char := range []rune(line) {
			if char == 'S' || char == 'a' {
				start = append(start, point.NewPoint2D(x, y))
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
	var values sort.IntSlice
	for _, s := range start {
		value, _ := findSummit(s, end, 0, grid, emptyMap)
		values = append(values, value)
	}
	sort.Sort(values)
	return values[0], nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Having chosen a better starting location, it now only took %d steps to reach the summit\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
