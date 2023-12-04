package day02

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	options = map[*regexp.Regexp]int{
		rRed:   12,
		rGreen: 13,
		rBlue:  14,
	}
)

func Part1Val(lines []string) (int, error) {
	var value int

	valid := func(round string, regex *regexp.Regexp, max int) bool {
		r := regex.FindStringSubmatch(round)
		if r == nil {
			return true
		}
		rVal, err := strconv.Atoi(r[1])
		if err != nil {
			return true
		}
		if rVal > max {
			return false
		}
		return true
	}

lines:
	for _, line := range lines {
		gameInfo, roundInfo, _ := strings.Cut(line, ":")

		var number int

		fmt.Sscanf(gameInfo, "Game %d", &number)

		for _, round := range strings.Split(roundInfo, ";") {
			for k, v := range options {
				if !valid(round, k, v) {
					continue lines
				}
			}
		}
		value += number
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
	fmt.Fprintf(w, "The total for the rocks was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
