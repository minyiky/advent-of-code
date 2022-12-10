package day09

import (
	"fmt"
	"io"
)

func Part1Val(lines []string) (int, error) {
	return simulateKnots(lines, 2)
}

func Part1(w io.Writer, lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The knotted tail of the rope covered %d of the squares in the gird\n", value)
	return nil
}
