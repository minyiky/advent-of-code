package day04

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	grid := make(map[Point]struct{})

	for j, line := range lines {
		for i, char := range line {
			if char == '@' {
				grid[Point{i, j}] = struct{}{}
			}
		}
	}

	setToCheck := grid
	for {
		newSetToCheck := make(map[Point]struct{}, len(grid))
		safeSet := make(map[Point]struct{}, len(setToCheck))
	setCheck:
		for point := range setToCheck {
			numBordering := 0
			for _, d := range AllDirs {
				if _, ok := grid[point.Add(d)]; ok {
					numBordering++
					if numBordering == 4 {
						continue setCheck
					}
				}
			}
			safeSet[point] = struct{}{}
		}

		// Add neighbors of safe points to check next
		for point := range safeSet {
			for _, d := range AllDirs {
				neighbor := point.Add(d)
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
