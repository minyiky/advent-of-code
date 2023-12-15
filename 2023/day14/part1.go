package day14

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part1Val(lines []string) (int, error) {
	var value int

	rounds := make(map[point.Point2D]bool)
	squares := make(map[point.Point2D]bool)

	lenLines := len(lines)

	for j, line := range lines {
		for i, char := range line {
			if char == 'O' {
				k := j
				for {
					if k == 0 ||
						rounds[point.NewPoint2D(i, k-1)] ||
						squares[point.NewPoint2D(i, k-1)] {
						rounds[point.NewPoint2D(i, k)] = true
						value += lenLines - k
						break
					}
					k--
				}
			} else if char == '#' {
				squares[point.NewPoint2D(i, j)] = true
			}
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
