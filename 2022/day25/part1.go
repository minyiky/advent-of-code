package day25

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

func Part1Val(lines []string) (string, error) {
	var value int

	for _, line := range lines {
		valList := []int{}
		for _, char := range line {
			switch char {
			case '-':
				valList = append([]int{-1}, valList...)
			case '=':
				valList = append([]int{-2}, valList...)
			default:
				val, err := strconv.Atoi(string(char))
				if err != nil {
					return "", err
				}
				valList = append([]int{val}, valList...)
			}
		}
		for i, val := range valList {
			tmpVal := val
			for j := 0; j < i; j++ {
				tmpVal *= 5
			}
			value += tmpVal
		}
	}

	SNAFU := ""
	for value > 0 {
		tmpVal := value % 5
		value /= 5

		switch tmpVal {
		case 4:
			SNAFU = "-" + SNAFU
			value += 1
		case 3:
			SNAFU = "=" + SNAFU
			value += 1
		default:
			SNAFU = strconv.Itoa(tmpVal) + SNAFU
		}
	}
	return SNAFU, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "You need to supply a SNAFU number of %s to bobs console\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
