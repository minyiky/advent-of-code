package day15

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string) (int, error) {
	var value = math.MaxInt

	costMap := make(map[point.Point2D]int)
	baseMap := make(map[point.Point2D]int)

	for i, line := range lines {
		for j, c := range line {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					baseMap[point.NewPoint2D(j+(k*len(line)), i+(l*len(lines)))] = (((int(c-'0') + k + l) - 1) % 9) + 1
					costMap[point.NewPoint2D(j+(k*len(line)), i+(l*len(lines)))] = math.MaxInt - 10000
				}
			}
		}
	}

	costMap[point.NewPoint2D(0, 0)] = 0

	var num int

	value = searchPriority(point.NewPoint2D(0, 0), point.NewPoint2D(len(lines[0])*5-1, len(lines)*5-1), baseMap, costMap, value, &num)
	// value = 0

	fmt.Println(num)

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
