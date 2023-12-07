package day07

import (
	"fmt"
	"io"
	"slices"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	hands := make([]Hand, len(lines))

	for i, line := range lines {
		hands[i] = NewJackHand(line)
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

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Your winnings with Jokers are: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
