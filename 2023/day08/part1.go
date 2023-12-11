package day08

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	nodes := make(map[string]Node)

	for _, line := range lines[2:] {
		matches := r.FindAllString(line, -1)
		node := Node{
			paths: map[rune]string{
				'L': matches[1],
				'R': matches[2],
			},
		}
		nodes[matches[0]] = node
	}

	i := 1

	node := nodes["AAA"]

loop:
	for {
		for _, dir := range lines[0] {
			next := node.paths[dir]
			if next == "ZZZ" {
				break loop
			}
			node = nodes[next]
			i++
		}
	}
	return i, nil
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
