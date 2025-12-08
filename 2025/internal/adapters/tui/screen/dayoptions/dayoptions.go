package dayoptions

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/navigation"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
)

var options = []string{
	"Run Part 1",
	"Run Part 2",
	"Run Both Parts",
	"Visualize Part 1",
	"Visualize Part 2",
}

type View struct {
	parentControls []navigation.Control
	dayNumber      int
	cursor         int
	SelectedOption int // Public field to communicate selected option to app
}

var _ screen.View = (*View)(nil)

func NewView(dayNumber int, parentControls ...navigation.Control) *View {
	return &View{
		parentControls: parentControls,
		dayNumber:      dayNumber,
		cursor:         0,
	}
}

func (v *View) Update(dayNumber int) {
	v.dayNumber = dayNumber
}

func (v *View) Render() string {
	s := fmt.Sprintf("Day %d - Select execution mode:\n\n", v.dayNumber)

	for i, option := range options {
		cursor := " "
		if i == v.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, option)
	}

	s += "\n\n"
	s += v.renderControls()
	return s
}

func (v *View) HandleMessage(msg tea.KeyMsg) *screen.ViewState {
	switch msg.String() {
	case "up", "k":
		if v.cursor > 0 {
			v.cursor--
		}

	case "down", "j":
		if v.cursor < len(options)-1 {
			v.cursor++
		}

	case "enter":
		// Set selected option for app to read
		v.SelectedOption = v.cursor
		// Navigate to running view
		nextView := screen.ViewRunning
		return &nextView
	}

	return nil
}

func (v *View) renderControls() string {
	controls := []navigation.Control{
		{
			Keys:        []string{"↑", "↓"},
			Instruction: "navigate",
		},
		{
			Keys:        []string{"Enter"},
			Instruction: "select",
		},
	}

	allControls := make([]string, 0, len(controls)+len(v.parentControls))
	for _, c := range controls {
		allControls = append(allControls, c.String())
	}
	for _, c := range v.parentControls {
		allControls = append(allControls, c.String())
	}
	return strings.Join(allControls, ", ")
}

func (v *View) Reset() {
	v.cursor = 0
}
