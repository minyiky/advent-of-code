package days

import (
	"github.com/minyiky/advent-of-code/2025/day02"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/visualization/rangepanels"
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

type Day02 struct {
	rangePanels *rangepanels.RangePanels
}

func (d *Day02) Part1Val(input []string) (int, error) {
	return day02.Part1Val(input)
}

func (d *Day02) Part2Val(input []string) (int, error) {
	return day02.Part2Val(input)
}

func (d *Day02) Part1ValWithViz(input []string, vizChan chan<- entities.VizFrame, speed float64) (int, error) {
	// Initialize range panels if needed
	if d.rangePanels == nil {
		d.rangePanels = rangepanels.New()
	}

	// Create a callback that renders the range panels and sends to channel
	callback := func(data map[string]interface{}) {
		d.renderAndSend(data, vizChan)
	}

	return day02.Part1ValWithViz(input, callback, speed)
}

func (d *Day02) Part2ValWithViz(input []string, vizChan chan<- entities.VizFrame, speed float64) (int, error) {
	// Initialize range panels if needed
	if d.rangePanels == nil {
		d.rangePanels = rangepanels.New()
	}

	// Create a callback that renders the range panels and sends to channel
	callback := func(data map[string]interface{}) {
		d.renderAndSend(data, vizChan)
	}

	return day02.Part2ValWithViz(input, callback, speed)
}

// renderAndSend extracts range panel data, renders it, and sends to channel
func (d *Day02) renderAndSend(data map[string]interface{}, vizChan chan<- entities.VizFrame) {
	// Extract range panel data with safe defaults
	currentRange := ""
	if cr, ok := data["currentRange"].(string); ok {
		currentRange = cr
	}

	rangeIndex := 0
	if ri, ok := data["rangeIndex"].(int); ok {
		rangeIndex = ri
	}

	totalRanges := 0
	if tr, ok := data["totalRanges"].(int); ok {
		totalRanges = tr
	}

	currentSum := 0
	if cs, ok := data["currentSum"].(int); ok {
		currentSum = cs
	}

	// Extract panels - convert from map[string]interface{} to map[string][]string
	panels := make(map[string][]string)
	if panelsData, ok := data["panels"].(map[string]interface{}); ok {
		for k, val := range panelsData {
			if ifaceSlice, ok := val.([]interface{}); ok {
				strSlice := make([]string, 0, len(ifaceSlice))
				for _, item := range ifaceSlice {
					if str, ok := item.(string); ok {
						strSlice = append(strSlice, str)
					}
				}
				if len(strSlice) > 0 {
					panels[k] = strSlice
				}
			}
		}
	}

	// Render the range panels (usecase owns the visualization)
	rendered := d.rangePanels.Render(currentRange, rangeIndex, totalRanges, panels, currentSum)

	// Send rendered output to channel (no extra delay needed for range panels)
	if vizChan != nil {
		vizChan <- entities.VizFrame{
			Rendered:              rendered,
			IsInstructionComplete: false, // Range panels don't need extra delay
		}
	}
}

func NewDay02() entities.VisualizableDay {
	d := &Day02{}
	return entities.NewBaseDay(d, 2)
}
