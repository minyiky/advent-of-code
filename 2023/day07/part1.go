package day07

import (
	"fmt"
	"io"
	"slices"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	hands := make([]Hand, len(lines))

	for i, line := range lines {
		hands[i] = NewHand(line)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.score != b.score {
			return a.score - b.score
		}
		if a.hand > b.hand {
			return 1
		}
		return -1
	})

	for i, hand := range hands {
		value += hand.bid * (i + 1)
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Your winnings are: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
