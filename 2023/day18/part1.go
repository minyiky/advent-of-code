package day18

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

var directions = map[string]point.Point2D{
	"U": point.NewPoint2D(0, 1),
	"D": point.NewPoint2D(0, -1),
	"L": point.NewPoint2D(-1, 0),
	"R": point.NewPoint2D(1, 0),
}

type limits struct {
	minX, maxX, minY, maxY int
}

func (l limits) Outside(p point.Point2D) bool {
	return p.X() < l.minX || p.X() > l.maxX || p.Y() < l.minY || p.Y() > l.maxY
}

func WayOut(pos point.Point2D, blocks, seen, outside map[point.Point2D]bool, limits limits) bool {
	if blocks[pos] {
		return false
	}

	if outside[pos] || limits.Outside(pos) {
		return true
	}

	seen[pos] = true

	for _, direction := range directions {
		newPoint := point.Add(pos, direction)
		if !seen[newPoint] {
			if WayOut(newPoint, blocks, seen, outside, limits) {
				return true
			}
		}
	}
	return false
}

func shoelace(vertices []point.Point2D) int {
	var area int

	for i := 0; i < len(vertices); i++ {
		j := (i + 1) % len(vertices)
		a, b := vertices[i], vertices[j]
		area += (a.X() * b.Y()) - (b.X() * a.Y()) + max(maths.Abs(a.X()-b.X()), maths.Abs(a.Y()-b.Y()))
	}

	return maths.Abs(area / 2)
}

func Part1Val(lines []string) (int, error) {
	var value int

	pos := point.NewPoint2D(0, 0)

	vertices := []point.Point2D{
		pos,
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		dir := directions[fields[0]]
		num, _ := strconv.Atoi(fields[1])

		value += num

		pos = point.Add(pos, point.NewPoint2D(dir.X()*num, dir.Y()*num))
		vertices = append(vertices, pos)
	}

	value += shoelace(vertices) + 1

	return value, nil
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
