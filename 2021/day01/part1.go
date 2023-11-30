package day01

import (
	"fmt"
	"io"
)

func Part1Val(lines []int) (int, error) {
	var numDeeper int
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			numDeeper++
		}
	}
	return numDeeper, nil
}

func Part1(w io.Writer, lines []int) error {
	val, err := Part1Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The depth measurement increased %d times\n", val)
	return nil
}
