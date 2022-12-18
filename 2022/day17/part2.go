package day17

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(line string) (int, error) {
	return QuickHeight(1000000000000, line), nil
}

func Part2(w io.Writer, line string) error {
	start := time.Now()
	value, err := Part2Val(line)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
