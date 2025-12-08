package tui

import "github.com/minyiky/advent-of-code/2025/internal/entities"

// DaysDiscoveredMsg is sent when days have been discovered
type DaysDiscoveredMsg struct {
	Days []int
	Err  error
}

// DayInfosLoadedMsg is sent when day infos are loaded from repository
type DayInfosLoadedMsg struct {
	Infos []*entities.DayInfo
}

// DayRunCompleteMsg is sent when a day run completes
type DayRunCompleteMsg struct {
	DayInfo *entities.DayInfo
	Err     error
}

// ProgressMsg is sent to update progress status
type ProgressMsg string

// VisualizationMsg is sent to update visualization state
type VisualizationMsg struct {
	Type string                 // Type of visualization (e.g., "dial")
	Data map[string]interface{} // Visualization-specific data
}

// ErrorMsg is sent when an error occurs
type ErrorMsg struct {
	Err error
}
