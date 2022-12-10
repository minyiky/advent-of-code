package day04

import (
	"fmt"
	"io"
)

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		var elf1S, elf2S, elf1E, elf2E int

		if _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &elf1S, &elf1E, &elf2S, &elf2E); err != nil {
			return 0, err
		}

		if !(elf1S > elf2E || elf2S > elf1E) {
			value += 1
		}

	}
	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "For %d of the pairs of elves, at least some level of overlap was found\n", value)
	return nil
}
