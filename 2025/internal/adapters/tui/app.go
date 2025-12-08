package tui

import (
	"context"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/navigation"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen/benchmark"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen/daylist"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen/dayoptions"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen/running"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui/screen/welcome"
	"github.com/minyiky/advent-of-code/2025/internal/entities"
	"github.com/minyiky/advent-of-code/2025/internal/usecases"
)

// App is the main TUI application model
type App struct {
	// Use cases
	runDay *usecases.RunDayUseCase

	// Repositories
	dayRepo usecases.DayRepository

	// Gateways
	dayLoader usecases.DayLoader

	// State
	days              []int
	dayInfos          map[int]*entities.DayInfo
	width             int
	height            int
	err               error
	loading           bool
	status            string
	progressSub       chan string
	vizSub            chan VisualizationMsg
	viewState         screen.ViewState
	outputBuffer      []string
	currentRunningDay int
	selectedDay       int

	// Visualization state
	renderedVisualization     string
	hasVisualization          bool
	speedIndex                atomic.Int32      // Speed index (0-10), interpretation is per-visualization
	speedIndexToActionsPerSec func(int32) int32 // Function to convert speed index to actions/sec (set per visualization)
	cancelViz                 context.CancelFunc // Function to cancel current visualization

	ViewHistory    []screen.ViewState
	CurrentView    screen.View
	ViewWelcome    *welcome.View    // Welcome screen
	ViewBenchmark  *benchmark.View  // Benchmark screen
	ViewDayList    *daylist.View    // Day list screen
	ViewDayOptions *dayoptions.View // Day options screen
	ViewRunning    *running.View    // Running/execution screen
}

// getActionsPerSecond calculates actions per second from speed index
func (a *App) getActionsPerSecond() int32 {
	i := a.speedIndex.Load()
	if a.speedIndexToActionsPerSec == nil {
		return i
	}
	return a.speedIndexToActionsPerSec(i)
}

// handleSpeedControl handles speed adjustment keys (shared across visualizations)
// Returns true if a speed control key was handled
func (a *App) handleSpeedControl(key string) bool {
	switch key {
	case "left", "h", "-":
		// Decrease speed index (min 0)
		current := a.speedIndex.Load()
		a.speedIndex.Store(max(0, current-1))
		return true

	case "right", "l", "+", "=":
		// Increase speed index (max 10)
		current := a.speedIndex.Load()
		a.speedIndex.Store(min(10, current+1))
		return true
	}
	return false
}

// NewApp creates a new TUI application
func NewApp(
	runDay *usecases.RunDayUseCase,
	dayRepo usecases.DayRepository,
	dayLoader usecases.DayLoader,
) *App {
	app := &App{
		runDay:            runDay,
		dayRepo:           dayRepo,
		dayLoader:         dayLoader,
		dayInfos:          make(map[int]*entities.DayInfo),
		viewState:         screen.ViewWelcome,
		outputBuffer:      make([]string, 0),

		// Initialise Views
		ViewHistory: []screen.ViewState{screen.ViewWelcome},
	}
	app.ViewWelcome = welcome.NewView(navigation.DefaultControls()...)
	app.ViewBenchmark = benchmark.NewView(app.days, app.dayInfos, navigation.DefaultControls()...)
	app.ViewDayList = daylist.NewView(app.days, app.dayInfos, navigation.DefaultControls()...)
	app.ViewDayOptions = dayoptions.NewView(app.selectedDay, navigation.DefaultControls()...)
	app.ViewRunning = running.NewView(app.currentRunningDay, app.getActionsPerSecond(), app.renderedVisualization, app.outputBuffer, app.loading, app.hasVisualization)
	app.speedIndex.Store(0) // Start at 10 actions/sec (10 * 2^0)
	return app
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	return a.discoverDaysCmd()
}

// Update handles messages and updates the model
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return a.handleKeyPress(msg)

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		return a, nil

	case DaysDiscoveredMsg:
		if msg.Err != nil {
			a.err = msg.Err
			return a, nil
		}
		a.days = msg.Days
		a.loading = false
		return a, a.loadDayInfosCmd()

	case DayInfosLoadedMsg:
		for _, info := range msg.Infos {
			a.dayInfos[info.Number] = info
		}
		return a, nil

	case DayRunCompleteMsg:
		if msg.Err != nil {
			a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Error: %v", msg.Err))
			a.loading = false
			return a, nil
		}
		a.dayInfos[msg.DayInfo.Number] = msg.DayInfo

		// Display results based on what was actually run
		var result string
		if msg.DayInfo.Part1Time > 0 && msg.DayInfo.Part2Time > 0 {
			result = fmt.Sprintf("Completed! Part1: %s, Part2: %s", msg.DayInfo.Part1Time, msg.DayInfo.Part2Time)
		} else if msg.DayInfo.Part1Time > 0 {
			result = fmt.Sprintf("Completed Part 1! Time: %s", msg.DayInfo.Part1Time)
		} else if msg.DayInfo.Part2Time > 0 {
			result = fmt.Sprintf("Completed Part 2! Time: %s", msg.DayInfo.Part2Time)
		}
		a.outputBuffer = append(a.outputBuffer, result)
		a.loading = false
		return a, nil

	case ProgressMsg:
		// Append progress message to output buffer
		a.outputBuffer = append(a.outputBuffer, string(msg))
		// Continue listening for more progress updates
		return a, waitForProgress(a.progressSub)

	case VisualizationMsg:
		// Update visualization state with rendered content
		if rendered, ok := msg.Data["rendered"].(string); ok {
			a.renderedVisualization = rendered
		}
		// Continue listening for more visualization updates
		return a, waitForVisualization(a.vizSub)

	case ErrorMsg:
		a.err = msg.Err
		a.loading = false
		a.status = ""
		return a, nil
	}

	return a, nil
}

// View renders the application
func (a *App) View() string {
	switch a.viewState {
	case screen.ViewWelcome:
		a.CurrentView = a.ViewWelcome
		return a.ViewWelcome.Render()
	case screen.ViewBenchmark:
		// Update benchmark view with latest data before rendering
		a.ViewBenchmark.Update(a.days, a.dayInfos)
		a.CurrentView = a.ViewBenchmark
		return a.ViewBenchmark.Render()
	case screen.ViewDayList:
		// Update day list view with latest data before rendering
		a.ViewDayList.Update(a.days, a.dayInfos)
		a.CurrentView = a.ViewDayList
		return a.ViewDayList.Render()
	case screen.ViewDayOptions:
		// Update day options view with latest data before rendering
		a.ViewDayOptions.Update(a.selectedDay)
		a.CurrentView = a.ViewDayOptions
		return a.ViewDayOptions.Render()
	case screen.ViewRunning:
		// Update running view with latest data before rendering
		a.ViewRunning.Update(a.currentRunningDay, a.getActionsPerSecond(), a.renderedVisualization, a.outputBuffer, a.loading, a.hasVisualization)
		a.CurrentView = a.ViewRunning
		return a.ViewRunning.Render()
	}

	if a.loading {
		return "Loading..."
	}

	return "Error: " + a.err.Error() + "\n\nPress q to quit"
}

// handleKeyPress handles keyboard input
func (a *App) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle quit globally
	if msg.String() == "q" || msg.String() == "ctrl+c" {
		return a, tea.Quit
	}

	if msg.String() == "esc" {
		// Cancel any running visualization
		if a.cancelViz != nil {
			a.cancelViz()
			a.cancelViz = nil
		}

		a.ViewHistory = a.ViewHistory[:len(a.ViewHistory)-1]
		if len(a.ViewHistory) == 0 {
			return a, tea.Quit
		}
		a.CurrentView.Reset()
		a.viewState = a.ViewHistory[len(a.ViewHistory)-1]
		return a, nil
	}

	// Handle view-specific key presses
	if newScr := a.CurrentView.HandleMessage(msg); newScr != nil {
		// View requested navigation, handle state updates
		switch *newScr {
		case screen.ViewDayOptions:
			// Coming from DayList, get selected day
			if a.viewState == screen.ViewDayList {
				a.selectedDay = a.ViewDayList.SelectedDay
			}
		case screen.ViewRunning:
			// Coming from DayOptions, get selected option and run
			if a.viewState == screen.ViewDayOptions {
				selectedOption := a.ViewDayOptions.SelectedOption
				a.ViewHistory = append(a.ViewHistory, *newScr)
				a.viewState = *newScr
				return a, a.runDayCmd(a.selectedDay, selectedOption)
			}
		}

		// Navigate to the new view
		a.ViewHistory = append(a.ViewHistory, *newScr)
		a.viewState = *newScr
		return a, nil
	}

	// Handle view-specific state updates (non-navigation keys)
	switch a.viewState {
	case screen.ViewRunning:
		// Try speed control (works even after completion)
		if a.handleSpeedControl(msg.String()) {
			return a, nil
		}

		// Ignore other keys while in running view
		return a, nil
	}

	return a, nil
}

// discoverDaysCmd discovers available days
func (a *App) discoverDaysCmd() tea.Cmd {
	return func() tea.Msg {
		days, err := a.dayLoader.DiscoverDays()
		return DaysDiscoveredMsg{Days: days, Err: err}
	}
}

// loadDayInfosCmd loads existing day info from repository
func (a *App) loadDayInfosCmd() tea.Cmd {
	return func() tea.Msg {
		infos, err := a.dayRepo.FindAll()
		if err != nil {
			return ErrorMsg{Err: err}
		}
		// Store the infos in a message to be processed
		return DayInfosLoadedMsg{Infos: infos}
	}
}

// runDayCmd runs a specific day with the selected part option
// partOption: 0 = Run Part 1, 1 = Run Part 2, 2 = Run Both, 3 = Visualize Part 1, 4 = Visualize Part 2
func (a *App) runDayCmd(dayNum int, partOption int) tea.Cmd {
	// Determine if this is a visualization run
	isVisualization := partOption >= 3

	// Cancel any previous visualization
	if a.cancelViz != nil {
		a.cancelViz()
		a.cancelViz = nil
	}

	// Switch to running view and reset state
	a.viewState = screen.ViewRunning
	a.loading = true
	a.outputBuffer = make([]string, 0)
	a.currentRunningDay = dayNum
	a.renderedVisualization = ""
	a.hasVisualization = isVisualization
	a.speedIndexToActionsPerSec = nil // Reset speed function to default

	// Set initial message based on part selection
	switch partOption {
	case 0:
		a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Running Day %d Part 1...", dayNum))
	case 1:
		a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Running Day %d Part 2...", dayNum))
	case 2:
		a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Running Day %d (Both Parts)...", dayNum))
	case 3:
		a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Visualizing Day %d Part 1...", dayNum))
	case 4:
		a.outputBuffer = append(a.outputBuffer, fmt.Sprintf("Visualizing Day %d Part 2...", dayNum))
	}

	// Create channels
	progressChan := make(chan string)
	vizChan := make(chan VisualizationMsg, 10)
	completionChan := make(chan DayRunCompleteMsg, 1)
	a.progressSub = progressChan
	a.vizSub = vizChan

	// Start the work in a goroutine
	go func() {
		inputPath := fmt.Sprintf("day%02d/input.txt", dayNum)

		// Progress callback
		progressFn := func(status string) {
			progressChan <- status
		}

		var dayInfo *entities.DayInfo
		var err error

		// Execute based on part selection
		if partOption == 2 {
			// Both parts - no visualization
			dayInfo, err = a.runDay.ExecuteBoth(dayNum, inputPath, progressFn)
		} else if isVisualization {
			// Visualization mode (options 3 or 4)
			part := 1
			if partOption == 4 {
				part = 2
			}

			// Configure speed function for visualization
			// Day 1 uses dial (faster), Day 2 uses range panels (slower)
			if dayNum == 1 {
				a.speedIndexToActionsPerSec = func(i int32) int32 {
					return 10 << i // Dial: 10 * 2^speedIndex
				}
			} else if dayNum == 2 {
				a.speedIndexToActionsPerSec = func(i int32) int32 {
					return 1 << i // Range panels: 1 * 2^speedIndex
				}
			}
			a.speedIndex.Store(0)

			// Create context for this visualization run
			ctx, cancel := context.WithCancel(context.Background())
			a.cancelViz = cancel

			// Create a channel for rendered visualization frames
			renderedChan := make(chan entities.VizFrame, 10)
			// Channel to signal when forwarding goroutine is done
			forwardingDone := make(chan struct{})

			// Start a goroutine to handle rendered output with delays
			go func() {
				defer close(forwardingDone)
				for {
					select {
					case <-ctx.Done():
						// Visualization cancelled, drain remaining messages
						for len(renderedChan) > 0 {
							<-renderedChan
						}
						return
					case frame, ok := <-renderedChan:
						if !ok {
							// Channel closed, stop processing
							return
						}
						// Send the rendered visualization to the UI
						select {
						case vizChan <- VisualizationMsg{Type: "rendered", Data: map[string]interface{}{
							"rendered": frame.Rendered,
						}}:
							// Message sent successfully
						case <-ctx.Done():
							// Cancelled while trying to send
							return
						}

						// Apply delay based on current speed
						actionsPerSec := a.getActionsPerSecond()
						if actionsPerSec == 0 {
							actionsPerSec = 1 // Safety: prevent divide by zero
						}
						delayMs := 1000.0 / float64(actionsPerSec)

						// Add extra delay if this completes an instruction (dial rotations)
						if frame.IsInstructionComplete {
							speedIdx := a.speedIndex.Load()
							if speedIdx > 0 {
								extraDelay := 200.0 / math.Sqrt(float64(speedIdx))
								delayMs += extraDelay
							} else {
								delayMs += 200.0
							}
						}

						select {
						case <-time.After(time.Duration(delayMs * float64(time.Millisecond))):
							// Delay complete
						case <-ctx.Done():
							// Cancelled during delay
							return
						}
					}
				}
			}()

			// Pass the channel to the usecase
			dayInfo, err = a.runDay.ExecuteWithVisualization(dayNum, part, inputPath, renderedChan, 1.0, progressFn)

			// Close the rendered channel to signal completion
			close(renderedChan)
			// Wait for the forwarding goroutine to finish processing
			<-forwardingDone
			// Now cancel the context to clean up
			cancel()
		} else {
			// Regular execution (options 0 or 1)
			part := 1
			if partOption == 1 {
				part = 2
			}
			dayInfo, err = a.runDay.Execute(dayNum, part, inputPath, progressFn)
		}

		// Close channels
		close(progressChan)
		close(vizChan)

		// Send completion message
		completionChan <- DayRunCompleteMsg{DayInfo: dayInfo, Err: err}
	}()

	// Return a batch command that waits for progress, viz, and completion
	return tea.Batch(
		waitForProgress(progressChan),
		waitForVisualization(vizChan),
		waitForCompletion(completionChan),
	)
}

// waitForProgress waits for the next progress update from a channel
func waitForProgress(sub chan string) tea.Cmd {
	return func() tea.Msg {
		status, ok := <-sub
		if !ok {
			// Channel closed, no more progress
			return nil
		}
		var msg ProgressMsg = ProgressMsg(status)
		return msg
	}
}

// waitForVisualization waits for the next visualization update from a channel
func waitForVisualization(vizChan chan VisualizationMsg) tea.Cmd {
	return func() tea.Msg {
		viz, ok := <-vizChan
		if !ok {
			// Channel closed, no more visualization
			return nil
		}
		return viz
	}
}

// waitForCompletion waits for the completion message from a channel
func waitForCompletion(completionChan chan DayRunCompleteMsg) tea.Cmd {
	return func() tea.Msg {
		return <-completionChan
	}
}
