package day16

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value int

	for i := range lines {
		cache := make(map[point.Point2D]map[point.Point2D]bool)
		move(point.NewPoint2D(0, i), right, lines, cache)
		value = max(value, len(cache))
	}

	for i := range lines {
		cache := make(map[point.Point2D]map[point.Point2D]bool)
		move(point.NewPoint2D(len(lines[0])-1, i), left, lines, cache)
		value = max(value, len(cache))
	}

	for i := range lines[0] {
		cache := make(map[point.Point2D]map[point.Point2D]bool)
		move(point.NewPoint2D(i, 0), down, lines, cache)
		value = max(value, len(cache))
	}

	for i := range lines[0] {
		cache := make(map[point.Point2D]map[point.Point2D]bool)
		move(point.NewPoint2D(i, len(lines)-1), up, lines, cache)
		value = max(value, len(cache))
	}

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
