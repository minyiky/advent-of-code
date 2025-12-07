package day04

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

func Part2Val(lines []string) (int, error) {
	var value int

	grid := make(map[point.Point]struct{})

	for j, line := range lines {
		for i, char := range line {
			if char == '@' {
				grid[point.Point{i, j}] = struct{}{}
			}
		}
	}

	setToCheck := grid
	for {
		newSetToCheck := make(map[point.Point]struct{}, len(grid))
		safeSet := make(map[point.Point]struct{}, len(setToCheck))
	setCheck:
		for p := range setToCheck {
			numBordering := 0
			for _, d := range point.AllDirs {
				if _, ok := grid[p.Add(d)]; ok {
					numBordering++
					if numBordering == 4 {
						continue setCheck
					}
				}
			}
			safeSet[p] = struct{}{}
		}

		// Add neighbors of safe points to check next
		for p := range safeSet {
			for _, d := range point.AllDirs {
				neighbor := p.Add(d)
				if _, ok := grid[neighbor]; ok {
					newSetToCheck[neighbor] = struct{}{}
				}
			}
		}

		value += len(safeSet)
		for point := range safeSet {
			delete(grid, point)
			delete(newSetToCheck, point)
		}

		if len(safeSet) == 0 {
			break
		}
		setToCheck = newSetToCheck
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
