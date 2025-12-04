package day03

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	codeLen := 12

	for _, line := range lines {
		v, err := findJoltage(line, codeLen)
		if err != nil {
			return 0, err
		}
		value += v
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
