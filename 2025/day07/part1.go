package day07

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

type traveller struct {
	maxHeight int
	splitters map[point.Point]struct{}
	traversed map[point.Point]struct{}
}

func (t *traveller) traverse(p point.Point) int {
	if _, ok := t.traversed[p]; ok {
		return 0
	}

	t.traversed[p] = struct{}{}

	if p.Y > t.maxHeight {
		return 0
	}

	nextP := p.Add(point.Down)

	if _, ok := t.splitters[nextP]; !ok {
		return t.traverse(nextP)
	}

	v := 1
	v += t.traverse(nextP.Add(point.Right))
	v += t.traverse(nextP.Add(point.Left))
	return v
}

func Part1Val(lines []string) (int, error) {
	var value int

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

	traversedPaths := make(map[point.Point]struct{})

	traveller := &traveller{
		maxHeight: maxHeight,
		splitters: splitters,
		traversed: traversedPaths,
	}

	value = traveller.traverse(start)
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
