package usecases

import (
	"fmt"

	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

// RunDayUseCase handles running a day's solution
type RunDayUseCase struct {
	dayRepo    DayRepository
	dayLoader  DayLoader
	fileReader FileReader
}

// NewRunDayUseCase creates a new RunDayUseCase
func NewRunDayUseCase(
	dayRepo DayRepository,
	dayLoader DayLoader,
	fileReader FileReader,
) *RunDayUseCase {
	return &RunDayUseCase{
		dayRepo:    dayRepo,
		dayLoader:  dayLoader,
		fileReader: fileReader,
	}
}

// Execute runs the specified day and part, updates timing info, and saves to repository
func (uc *RunDayUseCase) Execute(dayNumber int, part int, inputPath string, progressFn func(string)) (*entities.DayInfo, error) {
	if progressFn != nil {
		progressFn(fmt.Sprintf("Loading day %d implementation", dayNumber))
	}

	// Load the day implementation
	day, err := uc.dayLoader.Load(dayNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to load day %d: %w", dayNumber, err)
	}

	if progressFn != nil {
		progressFn(fmt.Sprintf("Reading input file: %s", inputPath))
	}

	// Read input file
	input, err := uc.fileReader.ReadLines(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	// Load or create day info
	dayInfo, err := uc.dayRepo.FindByNumber(dayNumber)
	if err != nil {
		// If not found, create new
		dayInfo = entities.NewDayInfo(dayNumber)
	}

	if progressFn != nil {
		progressFn(fmt.Sprintf("Running Part %d", part))
	}

	// Run the requested part and get timing + result
	duration, result, err := day.TimeAndRunPart(input, part)
	if err != nil {
		return nil, fmt.Errorf("failed to run part %d: %w", part, err)
	}

	// Update day info based on part
	if part == 1 {
		dayInfo.UpdatePart1(result, duration)
	} else {
		dayInfo.UpdatePart2(result, duration)
	}

	if progressFn != nil {
		progressFn("Saving results")
	}

	// Save updated info
	if err := uc.dayRepo.Save(dayInfo); err != nil {
		return nil, fmt.Errorf("failed to save day info: %w", err)
	}

	return dayInfo, nil
}

// ExecuteBoth runs both parts of a day efficiently (loading implementation and input once)
func (uc *RunDayUseCase) ExecuteBoth(dayNumber int, inputPath string, progressFn func(string)) (*entities.DayInfo, error) {
	if progressFn != nil {
		progressFn(fmt.Sprintf("Loading day %d implementation", dayNumber))
	}

	// Load the day implementation once
	day, err := uc.dayLoader.Load(dayNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to load day %d: %w", dayNumber, err)
	}

	if progressFn != nil {
		progressFn(fmt.Sprintf("Reading input file: %s", inputPath))
	}

	// Read input file once
	input, err := uc.fileReader.ReadLines(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	// Load or create day info
	dayInfo, err := uc.dayRepo.FindByNumber(dayNumber)
	if err != nil {
		// If not found, create new
		dayInfo = entities.NewDayInfo(dayNumber)
	}

	// Run part 1
	if progressFn != nil {
		progressFn("Running Part 1")
	}

	duration1, result1, err := day.TimeAndRunPart(input, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to run part 1: %w", err)
	}
	dayInfo.UpdatePart1(result1, duration1)

	// Run part 2
	if progressFn != nil {
		progressFn("Running Part 2")
	}

	duration2, result2, err := day.TimeAndRunPart(input, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to run part 2: %w", err)
	}
	dayInfo.UpdatePart2(result2, duration2)

	if progressFn != nil {
		progressFn("Saving results")
	}

	// Save updated info
	if err := uc.dayRepo.Save(dayInfo); err != nil {
		return nil, fmt.Errorf("failed to save day info: %w", err)
	}

	return dayInfo, nil
}

// ExecuteWithVisualization runs the specified day and part with visualization support
// Note: Visualization runs do NOT update benchmarks or save timing data
func (uc *RunDayUseCase) ExecuteWithVisualization(dayNumber int, part int, inputPath string, vizChan chan<- entities.VizFrame, speed float64, progressFn func(string)) (*entities.DayInfo, error) {
	// Load the day implementation
	day, err := uc.dayLoader.Load(dayNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to load day %d: %w", dayNumber, err)
	}

	// Read input file
	input, err := uc.fileReader.ReadLines(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	// Check if day supports visualization
	vizDay, supportsViz := day.(entities.VisualizableDay)
	if !supportsViz {
		// Fallback to regular execution
		return uc.Execute(dayNumber, part, inputPath, progressFn)
	}

	// Run with visualization (timing not saved)
	_, _, err = vizDay.RunPartWithVisualization(input, part, vizChan, speed)
	if err != nil {
		return nil, fmt.Errorf("failed to run part %d: %w", part, err)
	}

	// Load existing day info to return (don't create new or update)
	dayInfo, err := uc.dayRepo.FindByNumber(dayNumber)
	if err != nil {
		// If not found, create a minimal one just for display
		dayInfo = entities.NewDayInfo(dayNumber)
	}

	return dayInfo, nil
}
