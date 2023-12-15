package day15

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	line := lines[0]

	for _, part := range strings.Split(line, ",") {
		subval := 0
		for _, char := range part {
			subval += int(char)
			subval *= 17
			subval %= 256
		}
		value += subval
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
