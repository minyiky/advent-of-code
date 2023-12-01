package day01

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	value := 0

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		for i, word := range words {
			line = strings.Replace(line, word, fmt.Sprintf("%s%d%s", word[0:1], i, word[1:len(word)]), -1)
		}
		value += calculate(line)
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
	fmt.Fprintf(w, "Taking into account the spelled out digits as well the calibration value for the trebuchet was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
