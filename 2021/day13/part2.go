package day13

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (string, error) {
	points := make(map[point.Point2D]bool)

	var line string
	var i int

	for i, line = range lines {
		x, y, found := strings.Cut(line, ",")
		if !found {
			break
		}

		xNum, err := strconv.Atoi(x)
		if err != nil {
			return "", errors.New("bad input")
		}

		yNum, err := strconv.Atoi(y)
		if err != nil {
			return "", errors.New("bad input")
		}

		points[point.NewPoint2D(xNum, yNum)] = true

	}
	for _, line := range lines[i+1:] {
		section := ""
		fmt.Sscanf(line, "fold along %s", &section)
		dir, valStr, _ := strings.Cut(section, "=")
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return "", errors.New("bad input")
		}
		if dir == "x" {
			for p := range points {
				if p.X() > val {
					points[point.NewPoint2D(2*val-p.X(), p.Y())] = true
					delete(points, p)
				}

			}
		}
		if dir == "y" {
			for p := range points {
				if p.Y() > val {
					points[point.NewPoint2D(p.X(), 2*val-p.Y())] = true
					delete(points, p)
				}

			}
		}
	}

	xMax, yMax := 0, 0
	for p := range points {
		xMax = max(xMax, p.X())
		yMax = max(yMax, p.Y())
	}

	res := ""

	for y := 0; y <= yMax; y++ {
		if res != "" {
			res += "\n"
		}
		row := ""
		for x := 0; x <= xMax; x++ {
			if points[point.NewPoint2D(x, y)] {
				row += "#"
			} else {
				row += " "
			}
		}
		res += row
	}
	return res, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprint(w, value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
