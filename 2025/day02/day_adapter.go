package day02

import (
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

// day02 implements the concrete day logic
type day02 struct{}

// Part1Val implements part 1 logic
func (d *day02) Part1Val(input []string) (int, error) {
	return Part1Val(input)
}

// Part2Val implements part 2 logic
func (d *day02) Part2Val(input []string) (int, error) {
	return Part2Val(input)
}

// New creates a new day02 instance wrapped in BaseDay
func New() entities.Day {
	d := &day02{}
	return entities.NewBaseDay(d)
}
