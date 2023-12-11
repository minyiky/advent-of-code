package day08

import (
	"fmt"
	"io"
	"time"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part2Val(lines []string) (int, error) {
	nodes := make(map[string]Node)

	starts := make(map[string]struct{}, 0)

	for _, line := range lines[2:] {
		matches := r.FindAllString(line, -1)
		node := Node{
			paths: map[rune]string{
				'L': matches[1],
				'R': matches[2],
			},
		}
		nodes[matches[0]] = node
		if matches[0][2] == 'A' {
			starts[matches[0]] = struct{}{}
		}
	}

	i := 1

	lens := make([]int, 0, len(starts))

loop:
	for {
		for _, dir := range lines[0] {
			n := make(map[string]struct{})
			for name, _ := range starts {
				node := nodes[name]
				next := node.paths[dir]
				if next[2] == 'Z' {
					lens = append(lens, i)
					continue
				}
				n[next] = struct{}{}
			}
			if len(n) == 0 {
				break loop
			}
			starts = n
			i++
		}
	}
	value := LCM(lens[0], lens[1], lens[2:]...)
	return value, nil
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
