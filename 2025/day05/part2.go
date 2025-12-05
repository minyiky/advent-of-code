package day05

import (
	"fmt"
	"io"
	"slices"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	ranges := make([]Range, 0)

	for _, line := range lines {
		if line == "" {
			break
		}

		var a, b int
		if _, err := fmt.Sscanf(line, "%d-%d", &a, &b); err != nil {
			return 0, err
		}

		ranges = append(ranges, Range{Start: a, End: b})
	}

	slices.SortFunc(ranges, func(a, b Range) int { return a.Start - b.Start })

	combinedRanges := []Range{ranges[0]}

	for _, rng := range ranges[1:] {
		if rng.Start <= combinedRanges[len(combinedRanges)-1].End {
			if rng.End > combinedRanges[len(combinedRanges)-1].End {
				combinedRanges[len(combinedRanges)-1].End = rng.End
			}
		} else {
			combinedRanges = append(combinedRanges, rng)
		}
	}

	for _, rng := range combinedRanges {
		value += rng.End - rng.Start + 1
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
