package day05

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type seedInfo struct {
	start, length int
}

func reverse(input int, converters []converter) int {
	for _, converter := range converters {
		if input >= converter.target && input < converter.target+converter.length {
			return input - converter.target + converter.start
		}
	}
	return input
}

func checkSeeds(val int, seedInfos []seedInfo) bool {
	for _, info := range seedInfos {
		if val >= info.start && val < info.start+info.length {
			return true
		}
	}
	return false
}

func Part2Val(lines []string) (int, error) {
	seeds := make([]seedInfo, 0)

	vals := rNumber.FindAllString(lines[0], -1)

	for i := 0; i < len(vals)-1; i += 2 {
		start, _ := strconv.Atoi(vals[i])
		end, _ := strconv.Atoi(vals[i+1])
		seeds = append(seeds, seedInfo{
			start:  start,
			length: end,
		})
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

	i := 0
	for {
		seed := i
		for j := len(converterRoute) - 1; j >= 0; j-- {
			seed = reverse(seed, converterRoute[j])
		}
		if checkSeeds(seed, seeds) {
			return i, nil
		}
		i++
	}
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Using all the seeds, the lowest location found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
