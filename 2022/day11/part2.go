package day11

import (
	"fmt"
	"io"
)

func Part2Val(lines []string) (int, error) {
	return worryCalculator(lines, 1, 10000), nil
}

func Part2(w io.Writer, lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "These monkeys really are cheeky, after 10000 rounds of throwing items, these monkeys have managed to accrue a total monkey buisness of %d\n", value)
	return nil
}
