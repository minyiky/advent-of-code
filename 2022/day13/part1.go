package day13

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for i := 0; i < len(lines); i += 3 {
		left := lines[i]
		right := lines[i+1]
		left = StripEnds(left)
		right = StripEnds(right)
		valid, _ := isValid(left, right)
		if valid {
			value += (i / 3) + 1
		}
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
	fmt.Fprintf(w, "the total vaues for thre correct pairs of the distress signal was %d\n", value)
	fmt.Fprintf(w, "This took %.2fμs\n", float64(duration)/1e3)
	return nil
}
