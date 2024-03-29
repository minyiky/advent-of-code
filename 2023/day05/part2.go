package day05

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"time"
)

type seedInfo struct {
	start, length int
}

func translateRanges(input seedInfo, converters []converter) []seedInfo {
	outputs := make([]seedInfo, 0)
	for _, converter := range converters {
		if input.start >= converter.start && input.start < converter.start+converter.length {
			if input.start+input.length <= converter.start+converter.length {
				outputs = append(outputs, seedInfo{
					start:  input.start - converter.start + converter.target,
					length: input.length,
				})
				input.length = 0
				break
			}
			newLength := converter.start + converter.length - input.start
			outputs = append(outputs, seedInfo{
				start:  input.start - converter.start + converter.target,
				length: newLength,
			})
			input = seedInfo{
				start:  converter.start + converter.length,
				length: input.length - newLength,
			}
		}
	}
	if input.length > 0 {
		outputs = append(outputs, input)
	}
	return outputs
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

	inMap := true
	for _, line := range lines[3:] {
		if line == "" {
			inMap = false
			slices.SortFunc(converters, func(a, b converter) int {
				return a.start - b.start
			})
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

	slices.SortFunc(converters, func(a, b converter) int {
		return a.start - b.start
	})
	converterRoute = append(converterRoute, converters)

	finalSeeds := make([]seedInfo, 0, len(seeds))
	for _, seed := range seeds {
		newSeeds := []seedInfo{seed}
		for _, converters := range converterRoute {
			newNewSeeds := make([]seedInfo, 0, len(newSeeds))
			for _, s := range newSeeds {
				newNewSeeds = append(newNewSeeds, translateRanges(s, converters)...)
			}
			newSeeds = newNewSeeds
		}
		finalSeeds = append(finalSeeds, newSeeds...)
	}

	slices.SortFunc(finalSeeds, func(a, b seedInfo) int {
		return a.start - b.start
	})

	return finalSeeds[0].start, nil
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
