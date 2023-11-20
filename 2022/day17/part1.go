package day17

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func blockedSide(shape Shape, move point.Point2D, grid map[point.Point2D]bool) bool {
	for _, pos := range shape.points {
		newPos := point.Add(pos, move)
		if newPos.X() < 0 || newPos.X() > 6 {
			return true
		}

		if _, exists := grid[newPos]; exists {
			return true
		}
	}
	return false
}

func blockedDown(shape Shape, grid map[point.Point2D]bool) bool {
	down := point.NewPoint2D(0, -1)
	for _, pos := range shape.points {
		newPos := point.Add(pos, down)
		if _, exists := grid[newPos]; exists {
			return true
		}
	}
	return false
}

func Part1Val(line string) (int, error) {
	return QuickHeight(2022, line), nil
}

func Part1(w io.Writer, line string) error {
	start := time.Now()
	value, err := Part1Val(line)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "After 2022 blocks had fallen the tower had a total height of %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
