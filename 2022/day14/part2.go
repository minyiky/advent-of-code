package day14

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func sandFallRecursive(start point.Point2D, blocked map[point.Point2D]bool, yLim int) int {
	sand := start
	num := 1
	if sand.Y() == yLim {
		blocked[sand] = true
		return num
	}

	if _, ok := blocked[point.NewPoint2D(sand.X(), sand.Y()+1)]; !ok {
		num += sandFallRecursive(point.NewPoint2D(sand.X(), sand.Y()+1), blocked, yLim)
	}

	if _, ok := blocked[point.NewPoint2D(sand.X()-1, sand.Y()+1)]; !ok {
		num += sandFallRecursive(point.NewPoint2D(sand.X()-1, sand.Y()+1), blocked, yLim)
	}

	if _, ok := blocked[point.NewPoint2D(sand.X()+1, sand.Y()+1)]; !ok {
		num += sandFallRecursive(point.NewPoint2D(sand.X()+1, sand.Y()+1), blocked, yLim)
	}

	blocked[sand] = true
	return num
}

func Part2Val(lines []string) (int, error) {
	var value int

	blocked, yLim, err := createMap(lines)
	if err != nil {
		return 0, err
	}
	yLim += 1

	sandStart := point.NewPoint2D(500, 0)
	value = sandFallRecursive(sandStart, blocked, yLim)

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Now that you have seen there is a floor, you realiase the hole that the sand is pooring through will be covered after %d grains of sand have fallen\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
