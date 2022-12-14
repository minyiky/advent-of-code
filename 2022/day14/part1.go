package day14

import (
	"log"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines{
		_ = line
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The value found was: %d\n", value)
	return nil
}
