package day05

import (
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"time"
)

var rNumber = regexp.MustCompile(`([0-9]+)`)

type converter struct {
	start, length, target int
}

func translate(input int, converters []converter) int {
	for _, converter := range converters {
		if input >= converter.start && input < converter.start+converter.length {
			return input - converter.start + converter.target
		}
	}
	return input
}

func newConverter(line string) converter {
	vals := rNumber.FindAllString(line, -1)
	target, _ := strconv.Atoi(vals[0])
	start, _ := strconv.Atoi(vals[1])
	length, _ := strconv.Atoi(vals[2])
	return converter{
		start:  start,
		length: length,
		target: target,
	}
}

func Part1Val(lines []string) (int, error) {
	seeds := make([]int, 0)

	vals := rNumber.FindAllString(lines[0], -1)

	for _, val := range vals {
		seed, _ := strconv.Atoi(val)
		seeds = append(seeds, seed)
	}

	converterRoute := make([][]converter, 0)

	converters := make([]converter, 0)

	inMap := false
	for _, line := range lines[1:] {
		if line == "" {
			inMap = false
			converterRoute = append(converterRoute, converters)
			converters = make([]converter, 0)
			continue
		}

		if !inMap {
			inMap = true
			continue
		}

		converters = append(converters, newConverter(line))
	}

	converterRoute = append(converterRoute, converters)

	values := make([]int, len(seeds))

	for i, seed := range seeds {
		for _, converters := range converterRoute {
			seed = translate(seed, converters)
		}
		values[i] = seed
	}

	slices.Sort(values)

	return values[0], nil
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
