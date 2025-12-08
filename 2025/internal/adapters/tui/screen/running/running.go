package running

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
)

type View struct {
	dayNumber              int
	actionsPerSecond       int32
	renderedVisualization  string
	outputBuffer           []string
	loading                bool
	hasVisualization       bool
}

var _ screen.View = (*View)(nil)

func NewView(
	dayNumber int,
	actionsPerSecond int32,
	renderedVisualization string,
	outputBuffer []string,
	loading bool,
	hasVisualization bool,
) *View {
	return &View{
		dayNumber:             dayNumber,
		actionsPerSecond:      actionsPerSecond,
		renderedVisualization: renderedVisualization,
		outputBuffer:          outputBuffer,
		loading:               loading,
		hasVisualization:      hasVisualization,
	}
}

func (v *View) Update(
	dayNumber int,
	actionsPerSecond int32,
	renderedVisualization string,
	outputBuffer []string,
	loading bool,
	hasVisualization bool,
) {
	v.dayNumber = dayNumber
	v.actionsPerSecond = actionsPerSecond
	v.renderedVisualization = renderedVisualization
	v.outputBuffer = outputBuffer
	v.loading = loading
	v.hasVisualization = hasVisualization
}

func (v *View) Render() string {
	s := fmt.Sprintf("=== Running Day %d ===", v.dayNumber)
	s += v.renderSpeedIndicator()
	s += "\n\n"

	// Display pre-rendered visualization if available
	if v.renderedVisualization != "" {
		s += v.renderedVisualization
		s += "\n\n"
	}

	// Display all output messages
	for _, msg := range v.outputBuffer {
		s += msg + "\n"
	}

	// Show status/help text
	s += "\n"
	s += v.renderHelp()

	return s
}

func (v *View) HandleMessage(msg tea.KeyMsg) *screen.ViewState {
	// Speed controls are handled by the app, not by this view
	// This view only handles navigation away from the running view
	// (handled by parent via ESC)
	return nil
}

func (v *View) renderSpeedIndicator() string {
	if v.loading && v.hasVisualization {
		return fmt.Sprintf("  [Actions/sec: %d]", v.actionsPerSecond)
	}
	return ""
}

func (v *View) renderHelp() string {
	if v.loading {
		// Show controls during visualization
		if v.hasVisualization {
			return "\nRunning... (←/→ or +/- to adjust speed, ESC to cancel)"
		}
		return "\nRunning... (ESC to cancel)"
	}
	// After completion
	if v.hasVisualization {
		return "\n\nComplete! (←/→ or +/- to adjust speed for replay, ESC to return)"
	}
	return "\n\nComplete! (ESC to return to main menu)"
}

func (v *View) Reset() {
	// No special reset handling needed for running view
}
