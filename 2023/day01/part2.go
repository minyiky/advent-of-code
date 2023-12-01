package day01

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	return calculate(lines)
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Taking into account the spelled out digits as well the calibration value for the trebuchet was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
