package day01

import (
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

// day01 implements the concrete day logic with visualization support
type day01 struct{}

// Part1Val implements part 1 logic
func (d *day01) Part1Val(input []string) (int, error) {
	return Part1Val(input)
}

// Part2Val implements part 2 logic
func (d *day01) Part2Val(input []string) (int, error) {
	return Part2Val(input)
}

// Part1ValWithViz implements part 1 logic with visualization
func (d *day01) Part1ValWithViz(input []string, vizFactory entities.VizFnFactory, speed float64) (int, error) {
	return Part1ValWithViz(input, vizFactory, speed)
}

// Part2ValWithViz implements part 2 logic with visualization
func (d *day01) Part2ValWithViz(input []string, vizFactory entities.VizFnFactory, speed float64) (int, error) {
	return Part2ValWithViz(input, vizFactory, speed)
}

// New creates a new day01 instance wrapped in BaseDay
func New() entities.Day {
	d := &day01{}
	return entities.NewBaseDay(d)
}
