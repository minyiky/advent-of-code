package benchmark

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/navigation"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

type View struct {
	parentControls []navigation.Control
	days           []int
	dayInfos       map[int]*entities.DayInfo
}

var _ screen.View = (*View)(nil)

func NewView(days []int, dayInfos map[int]*entities.DayInfo, parentControls ...navigation.Control) *View {
	return &View{
		parentControls: parentControls,
		days:           days,
		dayInfos:       dayInfos,
	}
}

func (v *View) Update(days []int, dayInfos map[int]*entities.DayInfo) {
	v.days = days
	v.dayInfos = dayInfos
}

func (v *View) Render() string {
	s := "Advent of Code 2025 - Benchmarks\n\n"

	if len(v.days) == 0 {
		s += "No days found\n\n"
		s += v.renderControls()
		return s
	}

	// Find max time for scaling
	var maxTime time.Duration
	for _, dayNum := range v.days {
		if info, exists := v.dayInfos[dayNum]; exists {
			if info.Part1Time > maxTime {
				maxTime = info.Part1Time
			}
			if info.Part2Time > maxTime {
				maxTime = info.Part2Time
			}
		}
	}

	if maxTime == 0 {
		s += "No benchmark data available yet.\n"
		s += "Run some days first to see benchmark results!\n\n"
		s += v.renderControls()
		return s
	}

	// Graph settings
	const barWidth = 40
	const labelWidth = 15

	s += fmt.Sprintf("%-*s  Time\n", labelWidth, "Day/Part")
	s += fmt.Sprintf("%s\n", strings.Repeat("─", labelWidth+2+barWidth+15))

	for _, dayNum := range v.days {
		info, exists := v.dayInfos[dayNum]
		if !exists || (info.Part1Time == 0 && info.Part2Time == 0) {
			continue
		}

		// Part 1 bar
		if info.Part1Time > 0 {
			label := fmt.Sprintf("Day %d Part 1", dayNum)
			bars := int(float64(info.Part1Time) / float64(maxTime) * float64(barWidth))
			if bars == 0 && info.Part1Time > 0 {
				bars = 1 // At least show one bar if there's a time
			}
			s += fmt.Sprintf("%-*s  %s %s\n", labelWidth, label, strings.Repeat("█", bars), info.Part1Time)
		}

		// Part 2 bar
		if info.Part2Time > 0 {
			label := fmt.Sprintf("Day %d Part 2", dayNum)
			bars := int(float64(info.Part2Time) / float64(maxTime) * float64(barWidth))
			if bars == 0 && info.Part2Time > 0 {
				bars = 1 // At least show one bar if there's a time
			}
			s += fmt.Sprintf("%-*s  %s %s\n", labelWidth, label, strings.Repeat("█", bars), info.Part2Time)
		}
	}

	// Summary statistics
	s += "\n" + strings.Repeat("─", labelWidth+2+barWidth+15) + "\n"

	var totalTime time.Duration
	var count int
	for _, dayNum := range v.days {
		if info, exists := v.dayInfos[dayNum]; exists {
			if info.Part1Time > 0 {
				totalTime += info.Part1Time
				count++
			}
			if info.Part2Time > 0 {
				totalTime += info.Part2Time
				count++
			}
		}
	}

	if count > 0 {
		avgTime := totalTime / time.Duration(count)
		s += fmt.Sprintf("\nTotal Time: %s\n", totalTime)
		s += fmt.Sprintf("Average Time: %s\n", avgTime)
		s += fmt.Sprintf("Parts Completed: %d\n", count)
	}

	s += "\n\n"
	s += v.renderControls()

	return s
}

func (v *View) HandleMessage(msg tea.KeyMsg) *screen.ViewState {
	// No special key handling needed for benchmark view
	// All navigation is handled by parent controls
	return nil
}

func (v *View) Reset() {
	// No special reset handling needed for benchmark view
}

func (v *View) renderControls() string {
	allControls := make([]string, 0, len(v.parentControls))
	for _, c := range v.parentControls {
		allControls = append(allControls, c.String())
	}
	return strings.Join(allControls, ", ")
}
