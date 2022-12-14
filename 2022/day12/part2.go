package day12

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func Part2Val(lines []string) (int, error) {
	yLen, xLen := len(lines), len(lines[1])

	grid := make([][]rune, yLen)

	for i := range grid {
		grid[i] = make([]rune, xLen)
	}

	var start []aocutils.Vector
	var end aocutils.Vector
	for y, line := range lines {
		for x, char := range []rune(line) {
			if char == 'S' || char == 'a' {
				start = append(start, aocutils.NewVector(x, y))
				grid[y][x] = 'a'
				continue
			}
			if char == 'E' {
				end = aocutils.NewVector(x, y)
				grid[y][x] = 'z'
				continue
			}
			grid[y][x] = char
		}
	}

	emptyMap := make(map[aocutils.Vector]int)
	var values sort.IntSlice
	for _, s := range start {
		value, _ := findSummit(s, end, 0, grid, emptyMap)
		values = append(values, value)
	}
	sort.Sort(values)
	return values[0], nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Having chosen a better starting location, it now only took %d steps to reach the summit\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
