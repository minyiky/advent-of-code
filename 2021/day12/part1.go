package day12

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

func traverse(start string, small map[string][]string, large map[string][]string, visited map[string]bool, route []string) int {
	if start == "end" {
		return 1
	}
	if visited[start] {
		return 0
	}
	if _, ok := small[start]; ok {
		visited[start] = true
	}
	route = append(route, start)
	var count int
	for _, dest := range small[start] {
		newRoute := container.CopySlice(route)
		newVisited := container.CopyMap(visited)
		count += traverse(dest, small, large, newVisited, newRoute)
	}
	for _, dest := range large[start] {
		newRoute := container.CopySlice(route)
		newVisited := container.CopyMap(visited)
		count += traverse(dest, small, large, newVisited, newRoute)
	}
	return count
}

func Part1Val(lines []string) (int, error) {
	var value int

	small := make(map[string][]string)
	large := make(map[string][]string)

	for _, line := range lines {
		a, b, found := strings.Cut(line, "-")
		if !found {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		set := func(a, b string) {
			switch {
			case a == "end":
			case strings.ToLower(a) == a:
				if _, ok := small[a]; !ok {
					small[a] = make([]string, 0)
				}
				small[a] = append(small[a], b)
			default:
				if _, ok := large[a]; !ok {
					large[a] = make([]string, 0)
				}
				large[a] = append(large[a], b)
			}
		}
		set(a, b)
		set(b, a)

	}

	value = traverse("start", small, large, make(map[string]bool), make([]string, 0))

	return value, nil
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
