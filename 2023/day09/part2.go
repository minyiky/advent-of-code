package day09

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func findPrev(input []int) int {
	zeros := true
	diffs := make([]int, 0, len(input)-1)
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		diffs = append(diffs, diff)
		if diff != 0 {
			zeros = false
		}
	}
	if zeros {
		return input[0]
	}
	return input[0] - findPrev(diffs)
}

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		matches := rNUm.FindAllString(line, -1)
		nums := make([]int, len(matches))
		for i, match := range matches {
			nums[i], _ = strconv.Atoi(match)
		}
		value += findPrev(nums)
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
