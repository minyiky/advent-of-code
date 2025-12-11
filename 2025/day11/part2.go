package day11

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type DetailedVisit struct {
	count  int
	needed int
}

func Part2Val(lines []string) (int, error) {
	paths := make(map[string][]string)

	for _, line := range lines {
		start, targets, _ := strings.Cut(line, ":")
		paths[strings.TrimSpace(start)] = strings.Fields(strings.TrimSpace(targets))
	}

	visited := make(map[string]DetailedVisit)

	value := calc2("svr", paths, visited, 0)

	return value, nil
}

func calc2(node string, paths map[string][]string, visited map[string]DetailedVisit, neededNodes int) int {
	if node == "out" {
		if neededNodes == 2 {
			return 1
		} else {
			return 0
		}
	}

	if node == "fft" || node == "dac" {
		neededNodes++
	}

	prev, ok := visited[node]
	if ok && prev.needed == neededNodes {
		return prev.count
	}

	dv := DetailedVisit{count: 0, needed: neededNodes}

	for _, target := range paths[node] {
		dv.count += calc2(target, paths, visited, neededNodes)
	}

	visited[node] = dv

	return dv.count
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
