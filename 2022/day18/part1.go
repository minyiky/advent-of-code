package day18

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part1Val(lines []string) (int, error) {
	var value int

	grid := make(map[point.Point3D]bool)
	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		grid[point.NewPoint3D(x, y, z)] = true
	}

	for pos := range grid {
		for _, direction := range directions {
			if _, ok := grid[point.Add(pos, direction)]; !ok {
				value++
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
	fmt.Fprintf(w, "Your scan detescts that there are %d uncovered faces in the lava\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
