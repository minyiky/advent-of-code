package day04

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

func Part1Val(lines []string) (int, error) {
	grid := make(map[point.Point]struct{})

	for j, line := range lines {
		for i, char := range line {
			if char == '@' {
				grid[point.Point{i, j}] = struct{}{}
			}
		}
	}

	safeSet := make(map[point.Point]struct{})
	for p := range grid {
		numBordering := 0
		for _, d := range point.Cardinals {
			if _, ok := grid[p.Add(d)]; ok {
				numBordering++
			}
		}
		for _, d := range point.Diagonals {
			if _, ok := grid[p.Add(d)]; ok {
				numBordering++
			}
		}
		if numBordering < 4 {
			safeSet[p] = struct{}{}
		}
	}

	// newLines := make([]string, len(lines))
	// for j := 0; j < len(lines); j++ {
	// 	line := ""
	// 	for i := 0; i < len(lines[j]); i++ {
	// 		if _, ok := safeSet[point.Point{i, j}]; ok {
	// 			line += "x"
	// 		} else if _, ok := grid[point.Point{i, j}]; ok {
	// 			line += "@"
	// 		} else {
	// 			line += "."
	// 		}
	// 	}
	// 	newLines[j] = line
	// }
	// for _, line := range newLines {
	// 	fmt.Println(line)
	// }
	return len(safeSet), nil
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
