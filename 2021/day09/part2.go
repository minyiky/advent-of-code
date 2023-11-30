package day09

import (
	"fmt"
	"io"
	"slices"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type pos struct {
	height  rune
	visited bool
}

var (
	up    = point.NewPoint2D(0, 1)
	down  = point.NewPoint2D(0, -1)
	right = point.NewPoint2D(1, 0)
	left  = point.NewPoint2D(-1, 0)
)

func findBasin(heights map[point.Point2D]pos, p point.Point2D) int {
	height, ok := heights[p]

	if !ok || height.visited || height.height == '9' {
		return 0
	}

	height.visited = true
	heights[p] = height

	size := 1

	size += findBasin(heights, point.Add(p, up))
	size += findBasin(heights, point.Add(p, down))
	size += findBasin(heights, point.Add(p, right))
	size += findBasin(heights, point.Add(p, left))

	return size
}

func Part2Val(lines []string) (int, error) {
	var value int

	heights := make(map[point.Point2D]pos, 0)

	for i, line := range lines {
		for j, char := range line {
			heights[point.NewPoint2D(i, j)] = pos{char, false}
		}
	}

	basins := make([]int, 0)

	for p, h := range heights {
		if !h.visited {
			basins = append(basins, findBasin(heights, p))
		}
	}

	slices.Sort(basins)

	value = 1

	for i := 1; i < 4; i++ {
		value *= basins[len(basins)-i]
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
