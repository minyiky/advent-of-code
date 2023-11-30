package day15

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

var backwards = []point.Point2D{
	point.NewPoint2D(-1, 0),
	point.NewPoint2D(0, -1),
}

var allDirections = []point.Point2D{
	point.NewPoint2D(0, 1),
	point.NewPoint2D(1, 0),
	point.NewPoint2D(0, -1),
	point.NewPoint2D(-1, 0),
}

func closest(a, b point.Point2D) point.Point2D {
	if maths.Abs(a.X()-b.X()) > maths.Abs(a.Y()-b.Y()) {
		return point.NewPoint2D(1, 0)
	}
	return point.NewPoint2D(0, 1)
}

func Part1Val(lines []string) (int, error) {
	var value = math.MaxInt

	costMap := make(map[point.Point2D]int)
	baseMap := make(map[point.Point2D]int)

	for i, line := range lines {
		for j, c := range line {
			baseMap[point.NewPoint2D(j, i)] = int(c - '0')
			costMap[point.NewPoint2D(j, i)] = math.MaxInt - 10000
		}
	}

	costMap[point.NewPoint2D(0, 0)] = 0

	var num int

	value = searchPriority(point.NewPoint2D(0, 0), point.NewPoint2D(len(lines[0])-1, len(lines)-1), baseMap, costMap, value, &num)

	fmt.Println(num)

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
