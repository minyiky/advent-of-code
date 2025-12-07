package day06

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	numLen := len(lines) - 1

	nums := []int{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		s := ""
		for _, line := range lines[:numLen] {
			s += string(line[i])
		}

		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return 0, err
		}

		nums = append(nums, n)

		switch lines[numLen][i] {
		case '*':
			prod := 1
			for _, n := range nums {
				prod *= n
			}
			value += prod
			nums = []int{}
			i--
		case '+':
			for _, n := range nums {
				value += n
			}
			nums = []int{}
			i--
		}
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
