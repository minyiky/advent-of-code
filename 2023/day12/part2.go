package day12

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines{
		_ = line
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
