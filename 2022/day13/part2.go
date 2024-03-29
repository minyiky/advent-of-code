package day13

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

func Part2Val(lines []string) (int, error) {
	finalOrder := []string{
		"[[2]]",
		"[[6]]",
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		// TODO: Add binary sort
		var found bool
		tmpLine := StripEnds(line)
		for i, signal := range finalOrder {
			tmpSignal := StripEnds(signal)
			found, _ = isValid(tmpLine, tmpSignal)
			if found {
				finalOrder = append(finalOrder[:i], append([]string{line}, finalOrder[i:]...)...)
				break
			}
		}
		if !found {
			finalOrder = append(finalOrder, line)
		}
	}

	value := 1
	ofInterest := []string{
		"[[2]]",
		"[[6]]",
	}
	for i, signal := range finalOrder {
		if _, ok := container.SliceContains(ofInterest, signal); ok {
			value *= (i + 1)
		}
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
	fmt.Fprintf(w, "Now that you have sorted all of the packets, the diveiders have a combined position of %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
