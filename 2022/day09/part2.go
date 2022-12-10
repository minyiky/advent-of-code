package day09

import (
	"fmt"
	"io"
)

func Part2Val(lines []string) (int, error) {
	return simulateKnots(lines, 10)
}

func Part2(w io.Writer, lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "With the bridge broken and the end of of rope whipping around, the end of the rope covered %d squares in the grid\n", value)
	return nil
}
