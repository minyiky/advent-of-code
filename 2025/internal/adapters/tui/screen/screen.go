package screen

import tea "github.com/charmbracelet/bubbletea"

type ViewState string

const (
	ViewWelcome    ViewState = "Welcome"
	ViewDayList    ViewState = "Run Days"
	ViewDayOptions ViewState = "Day Options"
	ViewRunning    ViewState = "Running"
	ViewBenchmark  ViewState = "View Benchmark"
)

type View interface {
	HandleMessage(msg tea.KeyMsg) *ViewState
	Reset()
}
