package day02

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	calc := func(round string, regex *regexp.Regexp, max int) int {
		r := regex.FindStringSubmatch(round)
		if r == nil {
			return 0
		}
		rVal, err := strconv.Atoi(r[1])
		if err != nil {
			return 0
		}
		return rVal
	}

	for _, line := range lines {
		_, roundInfo, _ := strings.Cut(line, ":")

		number := 1

		vals := map[*regexp.Regexp]int{
			rRed:   0,
			rGreen: 0,
			rBlue:  0,
		}

		for _, round := range strings.Split(roundInfo, ";") {
			for k, v := range vals {
				if val := calc(round, k, v); val > v {
					vals[k] = val
				}
			}
		}

		for _, v := range vals {
			number *= v
		}

		value += number
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
	fmt.Fprintf(w, "Using the minimum number of rocks, the total for the rocks was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
