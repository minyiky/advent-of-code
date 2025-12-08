package welcome

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/navigation"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
)

const header = `
    _       _                 _            __    ____          _
   / \   __| |_   _____ _ __ | |_    ___  / _|  / ___|___   __| | ___
  / _ \ / _' \ \ / / _ \ '_ \| __|  / _ \| |_  | |   / _ \ / _' |/ _ \
 / ___ \ (_| |\ V /  __/ | | | |_  | (_) |  _| | |__| (_) | (_| |  __/
/_/   \_\__,_| \_/ \___|_| |_|\__|  \___/|_|    \____\___/ \__,_|\___|

                          2 0 2 5

`

var menuOptions = []screen.ViewState{
	screen.ViewDayList,
	screen.ViewBenchmark,
}

type View struct {
	parentControls []navigation.Control
	cursor         int
}

var _ screen.View = (*View)(nil)

func NewView(parentControls ...navigation.Control) *View {
	return &View{
		parentControls: parentControls,
	}
}

func (v *View) Render() string {
	s := header
	s += "\n\n"
	s += "Welcome to Advent of Code 2025!"
	s += "\n\n"

	for i, option := range menuOptions {
		cursor := " "
		if i == v.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, option)
	}

	s += "\n\n"

	controls := controls()
	allControls := make([]string, 0, len(controls)+len(v.parentControls))
	for _, c := range controls {
		allControls = append(allControls, c.String())
	}
	for _, c := range v.parentControls {
		allControls = append(allControls, c.String())
	}
	s += strings.Join(allControls, ", ")

	return s
}

func (v *View) HandleMessage(msg tea.KeyMsg) *screen.ViewState {
	switch msg.String() {
	case "up", "k":
		if v.cursor > 0 {
			v.cursor--
		}

	case "down", "j":
		if v.cursor < 1 {
			v.cursor++
		}

	case "enter":
		// Navigate based on selected menu option
		return &menuOptions[v.cursor]
	}

	return nil

}

func controls() []navigation.Control {
	return []navigation.Control{
		{
			Keys:        []string{"↑", "↓"},
			Instruction: "navigate",
		}, {
			Keys:        []string{"Enter"},
			Instruction: "run",
		},
	}
}

func (v *View) Reset() {
	v.cursor = 0
}
