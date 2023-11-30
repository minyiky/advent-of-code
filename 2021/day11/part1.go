package day11

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

var directions = []point.Point2D{
	point.NewPoint2D(1, 1),
	point.NewPoint2D(1, 0),
	point.NewPoint2D(1, -1),
	point.NewPoint2D(0, 1),
	point.NewPoint2D(0, -1),
	point.NewPoint2D(-1, 1),
	point.NewPoint2D(-1, 0),
	point.NewPoint2D(-1, -1),
}

func Part1Val(lines []string) (int, error) {
	var value int

	octos := make(map[point.Point2D]*octopus)

	for i, line := range lines {
		for j, c := range line {
			octos[point.NewPoint2D(i, j)] = &octopus{energy: int(c - '0')}
		}
	}

	for i := 0; i < 100; i++ {
		resets := make([]point.Point2D, 0)
		for _, o := range octos {
			o.increaseEnergy()
		}
		for {
			flashed := false
			for pos, o := range octos {
				if o.energy > 9 && !o.flashed {
					flashed = true
					o.flash()
					octos[pos] = o
					value++
					resets = append(resets, pos)
					for _, d := range directions {
						octo, ok := octos[point.Add(pos, d)]
						if !ok {
							continue
						}
						octo.increaseEnergy()
						octos[point.Add(pos, d)] = octo
					}
				}
			}
			if !flashed {
				break
			}
		}

		for _, pos := range resets {
			o := octos[pos]
			o.reset()
			octos[pos] = o
		}
	}

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
