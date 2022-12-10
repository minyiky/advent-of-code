package day03

import (
	"fmt"
	"io"
)

func Part1Val(lines []string) (int, error) {
	var score int

	for _, line := range lines {
		size := int(len(line) / 2)

		left, right := line[:size], line[size:]

		leftItems := getItems(left)

		match, err := getMatch(right, leftItems)
		if err != nil {
			return 0, err
		}

		val, err := getPriority(match)
		if err != nil {
			return 0, err
		}

		score += val
	}

	return score, nil
}

func Part1(w io.Writer, lines []string) error {
	score, err := Part1Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The shared items in the packs had a total priority of %d points\n", score)
	return nil
}
