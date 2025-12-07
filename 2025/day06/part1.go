package day06

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	if len(lines) < 4 {
		return 0, errors.New("not enough lines in the input")
	}

	nums := make([][]string, 0, len(lines)-1)
	ops := strings.Fields(lines[len(lines)-1])

	for _, line := range lines[:len(lines)-1] {
		nums = append(nums, strings.Fields(line))
	}

	for i := range ops {
		switch ops[i] {
		case "*":
			prod := 1
			for _, row := range nums {
				n, err := strconv.Atoi(row[i])
				if err != nil {
					return 0, err
				}
				prod *= n
			}
			value += prod
		case "+":
			for _, row := range nums {
				n, err := strconv.Atoi(row[i])
				if err != nil {
					return 0, err
				}
				value += n
			}
		default:
			return 0, fmt.Errorf("unexpected operator: %s", ops[i])
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
