package day03

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	commonBits := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				commonBits[i]++
			} else {
				commonBits[i]--
			}
		}
	}

	var gamma, epsilon int

	for i, bit := range commonBits {
		if bit > 0 {
			gamma |= 1 << (len(commonBits) - i - 1)
		} else {
			epsilon |= 1 << (len(commonBits) - i - 1)
		}
	}

	value := gamma * epsilon

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The power consumption is: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
