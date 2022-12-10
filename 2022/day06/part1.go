package day06

import (
	"fmt"
	"io"
)

func Part1Val(line string) (int, error) {
	return findSequence(line, 4)
}

func Part1(w io.Writer, line string) error {
	value, err := Part1Val(line)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "A marker sequence of 4 characters was found after receiving %d characters\n", value)
	return nil
}
