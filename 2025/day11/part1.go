package day11

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	paths := make(map[string][]string)

	for _, line := range lines {
		start, targets, _ := strings.Cut(line, ":")
		paths[strings.TrimSpace(start)] = strings.Fields(strings.TrimSpace(targets))
	}

	visited := make(map[string]int)

	value := calc("you", paths, visited)

	return value, nil
}

func calc(node string, paths map[string][]string, visited map[string]int) int {
	if node == "out" {
		return 1
	}

	if _, ok := visited[node]; ok {
		return visited[node]
	}

	for _, target := range paths[node] {
		visited[node] += calc(target, paths, visited)
	}

	return visited[node]
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
