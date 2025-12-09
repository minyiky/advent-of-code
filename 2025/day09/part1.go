package day09

import (
	"fmt"
	"io"
	"math"
	"slices"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

func Part1Val(lines []string) (int, error) {

	minX, minY := math.MaxInt, math.MaxInt

	points := make([]point.Point, 0, len(lines))

	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return 0, err
		}
		points = append(points, point.Point{X: x, Y: y})
		minX = min(minX, x)
		minY = min(minY, y)
	}

	for i, p := range points {
		p.X -= minX
		p.Y -= minY
		points[i] = p
	}

	slices.SortFunc(points, func(a, b point.Point) int {
		return b.Magnitude() - a.Magnitude()
	})

	currentMax := 0

	for i, p := range points {
		if p.Magnitude() < currentMax {
			break
		}

		for _, otherP := range points[i:] {
			size := p.Add(point.Point{X: 1, Y: 1}).Sub(otherP).Magnitude()
			currentMax = max(currentMax, size)
		}
	}

	return currentMax, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
