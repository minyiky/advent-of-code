package day14

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines{
		_ = line
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
