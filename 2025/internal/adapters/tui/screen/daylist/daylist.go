package daylist

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/navigation"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

type View struct {
	parentControls []navigation.Control
	days           []int
	dayInfos       map[int]*entities.DayInfo
	cursor         int
	SelectedDay    int // Public field to communicate selected day to app
}

var _ screen.View = (*View)(nil)

func NewView(days []int, dayInfos map[int]*entities.DayInfo, parentControls ...navigation.Control) *View {
	return &View{
		parentControls: parentControls,
		days:           days,
		dayInfos:       dayInfos,
		cursor:         0, // Start at first item
	}
}

func (v *View) Update(days []int, dayInfos map[int]*entities.DayInfo) {
	v.days = days
	v.dayInfos = dayInfos
}

func (v *View) Render() string {
	s := "Advent of Code 2025 - Select a Day\n\n"

	for i, dayNum := range v.days {
		cursor := " "
		if i == v.cursor {
			cursor = ">"
		}

		dayInfo, exists := v.dayInfos[dayNum]
		if exists {
			s += cursor + " Day " + string(rune(dayNum+'0')) + " - Part1: " + dayInfo.Part1Time.String() + ", Part2: " + dayInfo.Part2Time.String() + "\n"
		} else {
			s += cursor + " Day " + string(rune(dayNum+'0')) + " - Not run\n"
		}
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
		if v.cursor < len(v.days)-1 {
			v.cursor++
		}

	case "enter":
		if len(v.days) > 0 {
			// Set selected day for app to read
			v.SelectedDay = v.days[v.cursor]
			// Navigate to options view
			nextView := screen.ViewDayOptions
			return &nextView
		}
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
