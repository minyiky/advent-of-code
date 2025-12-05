package day05

import (
	"fmt"
	"io"
	"slices"
	"time"
)

type Range struct {
	Start int
	End   int
}

func Part1Val(lines []string) (int, error) {
	var value int
	line := ""
	i := 0

	ranges := make([]Range, 0)

	for i, line = range lines {
		if line == "" {
			break
		}

		var a, b int
		if _, err := fmt.Sscanf(line, "%d-%d", &a, &b); err != nil {
			return 0, err
		}

		ranges = append(ranges, Range{Start: a, End: b})
	}

	vals := make([]int, 0)

	for _, line := range lines[i+1:] {
		var a int
		if _, err := fmt.Sscanf(line, "%d", &a); err != nil {
			return 0, err
		}

		vals = append(vals, a)
	}

	slices.SortFunc(ranges, func(a, b Range) int { return a.Start - b.Start })

	slices.Sort(vals)

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

	combinedIdx := 0
	valIdx := 0
	for {
		if combinedIdx >= len(combinedRanges) || valIdx >= len(vals) {
			break
		}

		val := vals[valIdx]
		rng := combinedRanges[combinedIdx]

		// If the value is less than the start of the range, move to the next value
		if val < rng.Start {
			valIdx++
			continue
		}

		// If the value is greater than the end of the current range move to the next range
		if val > rng.End {
			combinedIdx++
			continue
		}

		// The value is in the range so increment the value and move to the next value
		value++
		valIdx++
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
