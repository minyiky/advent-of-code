package day01

import (
	"fmt"
	"io"
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
		value += int(code[0]-'0')*10 + int(code[len(code)-1]-'0')
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
