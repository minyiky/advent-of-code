package day07

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

type depthTraveller struct {
	maxHeight int
	splitters map[point.Point]struct{}
	traversed map[point.Point]int
}

func (t *depthTraveller) traverse(p point.Point) int {
	if val, ok := t.traversed[p]; ok {
		return val
	}

	if p.Y >= t.maxHeight {
		return 1
	}

	nextP := p.Add(point.Down)

	if _, ok := t.splitters[nextP]; !ok {
		v := t.traverse(nextP)
		t.traversed[p] = v
		return v
	}

	v := t.traverse(nextP.Add(point.Right))
	v += t.traverse(nextP.Add(point.Left))
	t.traversed[p] = v
	return v
}

type breadthTraveller struct {
	maxHeight int
	splitters map[point.Point]struct{}
	traversed map[point.Point]int
	queue     []point.Point
	head      int
	timelines int
}

func (t *breadthTraveller) traverse() {
	for t.head < len(t.queue) {
		p := t.queue[t.head]
		t.head++

		if p.Y >= t.maxHeight {
			t.timelines += t.traversed[p]
			continue
		}

		nextP := p.Add(point.Down)

		if _, ok := t.splitters[nextP]; !ok {
			if _, ok := t.traversed[nextP]; ok {
				t.traversed[nextP] += t.traversed[p]
			} else {
				t.traversed[nextP] = t.traversed[p]
				t.queue = append(t.queue, nextP)
			}
			continue
		}

		for _, splitP := range []point.Point{nextP.Add(point.Right), nextP.Add(point.Left)} {
			if _, ok := t.traversed[splitP]; ok {
				t.traversed[splitP] += t.traversed[p]
			} else {
				t.traversed[splitP] = t.traversed[p]
				t.queue = append(t.queue, splitP)
			}
		}
	}
}

func SolveBFS(maxHeight int, splitters map[point.Point]struct{}, start point.Point) int {
	traversedPaths := make(map[point.Point]int)

	breadthTraveller := &breadthTraveller{
		maxHeight: maxHeight,
		splitters: splitters,
		traversed: traversedPaths,
		queue:     []point.Point{start},
	}

	breadthTraveller.traversed[start] = 1
	breadthTraveller.traverse()

	return breadthTraveller.timelines
}

func SolveDFS(maxHeight int, splitters map[point.Point]struct{}, start point.Point) int {
	traversedPaths := make(map[point.Point]int)

	depthTraveller := &depthTraveller{
		maxHeight: maxHeight,
		splitters: splitters,
		traversed: traversedPaths,
	}

	return depthTraveller.traverse(start)
}

func ParseInput(lines []string) (int, map[point.Point]struct{}, point.Point) {
	maxHeight := len(lines)
	splitters := make(map[point.Point]struct{})
	var start point.Point

	for j, line := range lines {
		for i, char := range line {
			switch char {
			case 'S':
				start = point.Point{
					X: i,
					Y: j,
				}
			case '^':
				splitters[point.Point{
					X: i,
					Y: j,
				}] = struct{}{}
			}
		}
	}

	return maxHeight, splitters, start
}

func Part2ValBFS(lines []string) (int, error) {
	maxHeight, splitters, start := ParseInput(lines)
	return SolveBFS(maxHeight, splitters, start), nil
}

func Part2ValDFS(lines []string) (int, error) {
	maxHeight, splitters, start := ParseInput(lines)
	return SolveDFS(maxHeight, splitters, start), nil
}

func Part2Val(lines []string) (int, error) {
	return Part2ValDFS(lines)
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
