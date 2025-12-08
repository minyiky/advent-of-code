package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/gateways"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/repositories"
	"github.com/minyiky/advent-of-code/2025/internal/adapters/tui"
	"github.com/minyiky/advent-of-code/2025/internal/usecases"
	"github.com/minyiky/advent-of-code/2025/internal/usecases/days"
)

func main() {
	// Create day registry and register all days
	dayRegistry := gateways.NewDayRegistry(
		days.NewDay01(),
		days.NewDay02(),
		// Add more days here as they are implemented
	)

	// Create file reader
	fileReader := gateways.NewSimpleFileReader()

	// Create day repository for persistence
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get home directory: %w", err))
	}
	dayRepoPath := homeDir + "/.config/advent-of-code-tui/day_info.json"
	dayRepo, err := repositories.NewJSONDayRepository(dayRepoPath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create day repository: %w", err))
	}

	// Layer 2: Use Cases
	runDayUseCase := usecases.NewRunDayUseCase(dayRepo, dayRegistry, fileReader)

	// Layer 4: TUI App
	app := tui.NewApp(runDayUseCase, dayRepo, dayRegistry)

	// Run Bubbletea program
	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
