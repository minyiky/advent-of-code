package day11

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part1Val(lines []string) (int, error) {
	var value int

	spread := func(lines []string) []string {
		newLines := make([]string, 0, len(lines))

		for _, line := range lines {
			newLines = append(newLines, line)
			if !strings.Contains(line, "#") {
				newLines = append(newLines, line)
			}
		}
		return newLines
	}

	newLines := spread(lines)

	lines = make([]string, 0)

	for i := range newLines[0] {
		newLine := ""
		for _, line := range newLines {
			newLine += string(line[i])
		}
		lines = append(lines, newLine)
	}

	newLines = spread(lines)

	galaxies := make(map[point.Point2D]struct{})

	for y, line := range newLines {
		for x, char := range line {
			if char == '#' {
				galaxies[point.NewPoint2D(x, y)] = struct{}{}
			}
		}
	}

	for p1 := range galaxies {
		for p2 := range galaxies {
			value += point.ManhattanDistance(p1, p2)
		}
	}

	value /= 2

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
