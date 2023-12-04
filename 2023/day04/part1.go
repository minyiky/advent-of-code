package day04

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		_, line, _ = strings.Cut(line, ":")
		winners, numbers, _ := strings.Cut(line, "|")

		winSet := make(map[string]struct{})
		for _, winner := range strings.Fields(winners) {
			winSet[winner] = struct{}{}
		}

		power := -1
		for _, number := range strings.Fields(numbers) {
			if _, ok := winSet[number]; ok {
				power++
			}
		}
		if power >= 0 {
			value += int(math.Pow(2, float64(power)))
		}
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
