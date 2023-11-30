package day09

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	heights := make([][]byte, 0)

	for _, line := range lines {
		heights = append(heights, []byte(line))
	}

	for i, row := range heights {
		for j := range row {
			if (i == 0 || heights[i-1][j] > heights[i][j]) &&
				(j == 0 || heights[i][j-1] > heights[i][j]) &&
				(i == len(heights)-1 || heights[i+1][j] > heights[i][j]) &&
				(j == len(row)-1 || heights[i][j+1] > heights[i][j]) {
				value += int(heights[i][j] - '0' + 1)
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
