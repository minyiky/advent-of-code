package day11

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string, factor int) (int, error) {
	var value int

	rowCosts := make([]int, 0, len(lines))

	for _, line := range lines {
		if strings.Contains(line, "#") {
			rowCosts = append(rowCosts, 1)
		} else {
			rowCosts = append(rowCosts, factor)
		}
	}

	colCosts := make([]int, 0, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		line := ""
		for _, l := range lines {
			line += string(l[i])
		}
		if strings.Contains(line, "#") {
			colCosts = append(colCosts, 1)
		} else {
			colCosts = append(colCosts, factor)
		}
	}

	galaxies := make(map[point.Point2D]struct{})

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies[point.NewPoint2D(x, y)] = struct{}{}
			}
		}
	}

	for p1 := range galaxies {
		for p2 := range galaxies {
			getEnds := func(p1, p2 int) (int, int) {
				if p1 > p2 {
					return p2, p1
				}
				return p1, p2
			}

			startX, endX := getEnds(p1.X(), p2.X())
			startY, endY := getEnds(p1.Y(), p2.Y())

			row := 0

			for _, r := range rowCosts[startY:endY] {
				row += r
			}
			col := 0
			for _, c := range colCosts[startX:endX] {
				col += c
			}
			value += row + col
		}
	}

	value /= 2

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines, 1000000)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
