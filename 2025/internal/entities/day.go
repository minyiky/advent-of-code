package entities

import "time"

// Day represents a single Advent of Code day with its solution results
type Day interface {
	RunPart1(input []string) (int, error)
	RunPart2(input []string) (int, error)
	TimeAndRunPart(input []string, part int) (time.Duration, int, error)
	Number() int
}

// VizFrame represents a single frame of visualization with metadata
type VizFrame struct {
	Rendered              string // The rendered visualization content
	IsInstructionComplete bool   // True if this frame completes an instruction (for extra delay)
}

// VisualizableDay is an optional interface for days that support visualization
type VisualizableDay interface {
	Day
	RunPartWithVisualization(input []string, part int, vizChan chan<- VizFrame, speed float64) (time.Duration, int, error)
}
