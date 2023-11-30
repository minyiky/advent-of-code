package day05

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value int

	cells := make(map[point.Point2D]int)

	for _, line := range lines {
		var x1, x2, y1, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		calcMod := func(a, b int) int {
			if a > b {
				return -1
			} else if a < b {
				return 1
			}
			return 0
		}

		xMod := calcMod(x1, x2)
		yMod := calcMod(y1, y2)

		x := x1
		y := y1

		for {
			_, ok := cells[point.NewPoint2D(x, y)]
			if !ok {
				cells[point.NewPoint2D(x, y)] = 0
			}
			cells[point.NewPoint2D(x, y)]++

			y += yMod
			x += xMod
			if x == x2+xMod && y == y2+yMod {
				break
			}
		}
	}

	for _, v := range cells {
		if v > 1 {
			value++
		}
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
