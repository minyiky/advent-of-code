package day11

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value int

	octos := make(map[point.Point2D]*octopus)

	for i, line := range lines {
		for j, c := range line {
			octos[point.NewPoint2D(i, j)] = &octopus{energy: int(c - '0')}
		}
	}

	i := 1
	for {
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

		if len(resets) == 100 {
			break
		}

		for _, pos := range resets {
			o := octos[pos]
			o.reset()
			octos[pos] = o
		}
		i++
	}

	return i, nil
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
