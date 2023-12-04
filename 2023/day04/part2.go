package day04

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	scoreCards := make([]int, len(lines))
	lenLine := len(lines)

	for i := range lines {
		scoreCards[i] = 1
	}

	for i, line := range lines {
		_, line, _ = strings.Cut(line, ":")
		winners, numbers, _ := strings.Cut(line, "|")

		winSet := make(map[string]struct{})
		for _, winner := range strings.Fields(winners) {
			winSet[winner] = struct{}{}
		}

		numWinners := 0
		for _, number := range strings.Fields(numbers) {
			if _, ok := winSet[number]; ok {
				numWinners++
			}
		}

		for j := i + 1; j < min(i+1+numWinners, lenLine); j++ {
			scoreCards[j] += scoreCards[i]
		}
	}

	for _, scoreCard := range scoreCards {
		value += scoreCard
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
	fmt.Fprintf(w, "Using the updated scoring, the scratch cards scored: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
