package day17

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value int
	grid := make(map[point.Point2D]int)

	for j, line := range lines {
		for i, char := range line {
			grid[point.NewPoint2D(i, j)] = int(char - '0')
		}
	}

	start := point.NewPoint2D(0, 0)
	goal := point.NewPoint2D(len(lines[0])-1, len(lines)-1)

	value = run(start, goal, grid, 4, 10)

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
