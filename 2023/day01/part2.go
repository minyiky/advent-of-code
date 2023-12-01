package day01

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	keyVals := map[string]int{
		"0":     0,
		"zero":  0,
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	for _, line := range lines {
		lowVal, lowIndex := 0, math.MaxInt
		highVal, highIndex := 0, -1

		for k, v := range keyVals {
			index := strings.Index(line, k)
			if index < 0 {
				continue
			}
			if index < lowIndex {
				lowVal, lowIndex = v, index
			}

			index = strings.LastIndex(line, k)
			if index > highIndex {
				highVal, highIndex = v, index
			}
		}
		value += lowVal*10 + highVal
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
