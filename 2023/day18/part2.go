package day18

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

var newDirections = map[string]point.Point2D{
	"0": point.NewPoint2D(1, 0),
	"1": point.NewPoint2D(0, -1),
	"2": point.NewPoint2D(-1, 0),
	"3": point.NewPoint2D(0, 1),
}

func Part2Val(lines []string) (int, error) {
	var value int

	pos := point.NewPoint2D(0, 0)

	vertices := []point.Point2D{
		pos,
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		dir := newDirections[fields[2][len(fields[2])-2:len(fields[2])-1]]
		num, _ := strconv.ParseInt(fields[2][2:len(fields[2])-2], 16, 0)

		value += int(num)

		pos = point.Add(pos, point.NewPoint2D(dir.X()*int(num), dir.Y()*int(num)))
		vertices = append(vertices, pos)
	}

	value += shoelace(vertices) + 1

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
