package days

import (
	"github.com/minyiky/advent-of-code/2025/day01"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/visualization/dial"
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

type Day01 struct {
	dial         *dial.Dial
	instructions []string
}

func (d *Day01) Part1Val(input []string) (int, error) {
	return day01.Part1Val(input)
}

func (d *Day01) Part2Val(input []string) (int, error) {
	return day01.Part2Val(input)
}

func (d *Day01) Part1ValWithViz(input []string, vizChan chan<- entities.VizFrame, speed float64) (int, error) {
	// Initialize dial if needed
	if d.dial == nil {
		d.dial = dial.New()
	}
	d.instructions = input

	// Create a callback that renders the dial and sends to channel
	callback := func(data map[string]interface{}) {
		d.renderAndSend(data, vizChan)
	}

	return day01.Part1ValWithViz(input, callback, speed)
}

func (d *Day01) Part2ValWithViz(input []string, vizChan chan<- entities.VizFrame, speed float64) (int, error) {
	// Initialize dial if needed
	if d.dial == nil {
		d.dial = dial.New()
	}
	d.instructions = input

	// Create a callback that renders the dial and sends to channel
	callback := func(data map[string]interface{}) {
		d.renderAndSend(data, vizChan)
	}

	return day01.Part2ValWithViz(input, callback, speed)
}

// renderAndSend extracts dial data, renders it, and sends to channel
func (d *Day01) renderAndSend(data map[string]interface{}, vizChan chan<- entities.VizFrame) {
	// Extract dial-specific data
	position, _ := data["position"].(int)
	counter, _ := data["counter"].(int)
	maxValue, ok := data["maxValue"].(int)
	if !ok {
		maxValue = 100
	}
	finalPosition, ok := data["finalPosition"].(int)
	if !ok {
		finalPosition = position
	}
	instructionIndex, _ := data["instructionIndex"].(int)

	// Render the dial (usecase owns the visualization)
	rendered := d.dial.Render(position, counter, maxValue, finalPosition, d.instructions, instructionIndex)

	// Send rendered output to channel with metadata
	if vizChan != nil {
		vizChan <- entities.VizFrame{
			Rendered:              rendered,
			IsInstructionComplete: position == finalPosition,
		}
	}
}

func NewDay01() entities.VisualizableDay {
	d := &Day01{}
	return entities.NewBaseDay(d, 1)
}
