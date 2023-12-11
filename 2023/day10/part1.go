package day10

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type pipe struct {
	moves map[point.Point2D]point.Point2D
}

func (p pipe) move(m point.Point2D) (point.Point2D, error) {
	new, ok := p.moves[m]
	if !ok {
		// fmt.Printf("no move found for %v\n", m)
		return m, fmt.Errorf("no move found for %v", m)
	}
	return new, nil
}

var pipeTypes = map[rune]pipe{
	'|': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(0, 1):  point.NewPoint2D(0, 1),
			point.NewPoint2D(0, -1): point.NewPoint2D(0, -1),
		},
	},
	'-': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(1, 0):  point.NewPoint2D(1, 0),
			point.NewPoint2D(-1, 0): point.NewPoint2D(-1, 0),
		},
	},
	'L': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(0, -1): point.NewPoint2D(1, 0),
			point.NewPoint2D(-1, 0): point.NewPoint2D(0, 1),
		},
	},
	'J': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(1, 0):  point.NewPoint2D(0, 1),
			point.NewPoint2D(0, -1): point.NewPoint2D(-1, 0),
		},
	},
	'7': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(1, 0): point.NewPoint2D(0, -1),
			point.NewPoint2D(0, 1): point.NewPoint2D(-1, 0),
		},
	},
	'F': pipe{
		map[point.Point2D]point.Point2D{
			point.NewPoint2D(0, 1):  point.NewPoint2D(1, 0),
			point.NewPoint2D(-1, 0): point.NewPoint2D(0, -1),
		},
	},
}

var directions = []point.Point2D{
	point.NewPoint2D(0, 1),
	point.NewPoint2D(1, 0),
	point.NewPoint2D(0, -1),
	point.NewPoint2D(-1, 0),
}

func Part1Val(lines []string) (int, error) {
	var value int

	pipes := make(map[point.Point2D]pipe)
	var start point.Point2D

	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				start = point.NewPoint2D(x, -y)
				continue
			}
			if char == '.' {
				continue
			}
			pipes[point.NewPoint2D(x, -y)] = pipeTypes[char]
		}
	}

	i := 1

	pos := start
	var dir point.Point2D

	for _, d := range directions {
		pos = point.Add(start, d)
		// fmt.Printf("testing %c at %d:%d\n", lines[-pos.Y()][pos.X()], -pos.Y(), pos.X())
		var err error
		if dir, err = pipes[pos].move(d); err == nil {
			break
		}
	}

	for pos != start {
		i++
		pos = point.Add(pos, dir)
		dir, _ = pipes[pos].move(dir)
	}

	value = i / 2

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
