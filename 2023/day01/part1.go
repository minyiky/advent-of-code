package day01

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	codes := make([][]rune, len(lines))

	for i, line := range lines {
		for _, char := range line {
			if char-'0' >= 0 && char-'0' <= 9 {
				codes[i] = append(codes[i], char)
			}
		}
	}

	for _, code := range codes {
		numStr := fmt.Sprintf("%c%c", code[0], code[len(code)-1])
		val, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}
		value += val
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
	fmt.Fprintf(w, "The calibration value for the trebuchet was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
