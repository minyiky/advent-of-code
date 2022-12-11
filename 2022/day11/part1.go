package day11

import (
	"fmt"
	"io"
)

func Part1Val(lines []string) (int, error) {
	return worryCalculator(lines, 3, 20), nil
}

func Part1(w io.Writer, lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "There has been a total monkey buisness of %d occuring with your items\n", value)
	return nil
}
