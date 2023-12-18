package day17

import (
	"container/heap"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *PQ[T]) Push(x any) { *q = append(*q, x.(pqi[T])) }

func (q *PQ[_]) Pop() (x any) {
	x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return x
}

func (q *PQ[T]) GPush(v T, p int) { heap.Push(q, pqi[T]{v, p}) }

func (q *PQ[T]) GPop() (T, int) { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type State struct {
	Pos point.Point2D
	Dir point.Point2D
}

func run(
	start, end point.Point2D,
	grid map[point.Point2D]int,
	min, max int) int {
	queue := make(PQ[State], 0)
	seen := make(map[State]bool)

	queue.GPush(State{start, point.NewPoint2D(1, 0)}, 0)
	queue.GPush(State{start, point.NewPoint2D(0, 1)}, 0)

	for len(queue) > 0 {
		state, cost := queue.GPop()

		if state.Pos == end {
			return cost
		}
		if _, ok := seen[state]; ok && state.Pos != start {
			continue
		}
		seen[state] = true

		for i := -max; i <= max; i++ {
			newPoint := point.Add(state.Pos, point.NewPoint2D(state.Dir.X()*i, state.Dir.Y()*i))
			if _, ok := grid[newPoint]; !ok || i > -min && i < min {
				continue
			}
			c, s := 0, int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				newCost := grid[point.Add(state.Pos, point.NewPoint2D(state.Dir.X()*j, state.Dir.Y()*j))]
				c += newCost
			}
			queue.GPush(
				State{
					newPoint,
					point.NewPoint2D(state.Dir.Y(), state.Dir.X()),
				},
				cost+c)
		}
	}
	return -1
}

func Part1Val(lines []string) (int, error) {
	var value int
	grid := make(map[point.Point2D]int)

	for j, line := range lines {
		for i, char := range line {
			grid[point.NewPoint2D(i, j)] = int(char - '0')
		}
	}

	start := point.NewPoint2D(0, 0)
	goal := point.NewPoint2D(len(lines[0])-1, len(lines)-1)

	value = run(start, goal, grid, 1, 3)

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
