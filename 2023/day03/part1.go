package day03

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

	type values struct {
		p []point.Point2D
		v int
	}

	symbolMap := make(map[point.Point2D]struct{})
	valueMap := make([]values, 0)

	for i, line := range lines {
		v := values{
			p: make([]point.Point2D, 0),
			v: 0,
		}
		for j, c := range line {
			if c < '0' || c > '9' {
				if c != '.' {
					symbolMap[point.NewPoint2D(i, j)] = struct{}{}
				}
				if v.v != 0 {
					valueMap = append(valueMap, v)
					v = values{
						p: make([]point.Point2D, 0),
						v: 0,
					}
				}
				continue
			}
			v.v = v.v*10 + int(c-'0')
			v.p = append(v.p, point.NewPoint2D(i, j))
			if j == len(line)-1 {
				valueMap = append(valueMap, v)
				v = values{
					p: make([]point.Point2D, 0),
					v: 0,
				}
			}
		}
	}

values:
	for _, v := range valueMap {
		for _, d := range directions {
			for _, p := range v.p {
				if _, ok := symbolMap[point.Add(p, d)]; ok {
					value += v.v
					continue values
				}
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
