package entities

import "time"

// DayInfo stores metadata and timing information about a day's solution
type DayInfo struct {
	Number      int           `json:"number"`
	Part1Result int           `json:"part1_result"`
	Part2Result int           `json:"part2_result"`
	Part1Time   time.Duration `json:"part1_time"`
	Part2Time   time.Duration `json:"part2_time"`
	LastRun     time.Time     `json:"last_run"`
}

// NewDayInfo creates a new DayInfo for the given day number
func NewDayInfo(number int) *DayInfo {
	return &DayInfo{
		Number: number,
	}
}

// IsComplete returns true if both parts have been run
func (di *DayInfo) IsComplete() bool {
	return di.Part1Time > 0 && di.Part2Time > 0
}

// UpdatePart1 updates Part 1 result and timing
func (di *DayInfo) UpdatePart1(result int, duration time.Duration) {
	di.Part1Result = result
	di.Part1Time = duration
	di.LastRun = time.Now()
}

// UpdatePart2 updates Part 2 result and timing
func (di *DayInfo) UpdatePart2(result int, duration time.Duration) {
	di.Part2Result = result
	di.Part2Time = duration
	di.LastRun = time.Now()
}

// TotalTime returns the combined time for both parts
func (di *DayInfo) TotalTime() time.Duration {
	return di.Part1Time + di.Part2Time
}
