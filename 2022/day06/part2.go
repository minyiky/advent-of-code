package day06

import (
	"fmt"
	"io"
)

func Part2Val(line string) (int, error) {
	return findSequence(line, 14)
}

func Part2(w io.Writer, line string) error {
	value, err := Part2Val(line)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "A marker sequence of 14 characters was found after receiving %d characters\n", value)
	return nil
}
