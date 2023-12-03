package day03

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value int

	type values struct {
		p []point.Point2D
		v int
	}

	symbolMap := make(map[point.Point2D][]int)
	valueMap := make([]values, 0)

	for i, line := range lines {
		v := values{
			p: make([]point.Point2D, 0),
			v: 0,
		}
		for j, c := range line {
			if c < '0' || c > '9' {
				if c == '*' {
					symbolMap[point.NewPoint2D(i, j)] = make([]int, 0)
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

	for _, v := range valueMap {
		ignore := make(map[point.Point2D]bool)
		for _, d := range directions {
			for _, p := range v.p {
				newPoint := point.Add(p, d)
				if _, ok := symbolMap[newPoint]; ok {
					if _, ok := ignore[newPoint]; !ok {
						ignore[newPoint] = true
						symbolMap[newPoint] = append(symbolMap[newPoint], v.v)
					}
				}
			}
		}
	}

	for _, v := range symbolMap {
		if len(v) == 2 {
			value += v[0] * v[1]
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
