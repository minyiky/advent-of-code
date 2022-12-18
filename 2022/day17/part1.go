package day17

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func blockedSide(shape Shape, move aocutils.Vector, grid map[aocutils.Vector]bool) bool {
	for _, pos := range shape.points {
		newPos := pos.Add(move)
		if newPos.X < 0 || newPos.X > 6 {
			return true
		}

		if _, exists := grid[newPos]; exists {
			return true
		}
	}
	return false
}

func blockedDown(shape Shape, grid map[aocutils.Vector]bool) bool {
	down := aocutils.NewVector(0, -1)
	for _, pos := range shape.points {
		newPos := pos.Add(down)
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
