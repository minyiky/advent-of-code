package day12

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		cogPart, numPart, _ := strings.Cut(line, " ")

		numbersStr := rNum.FindAllString(numPart, -1)
		numbers := make([]int, len(numbersStr))
		for j, numberStr := range numbersStr {
			numbers[j], _ = strconv.Atoi(numberStr)
		}

		cache := make(map[Key]int)
		v := dp(cogPart, numbers, 0, 0, 0, cache)
		value += v
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
