package usecases

import "github.com/minyiky/advent-of-code/2025/internal/entities"

// DayRepository defines the interface for day metadata persistence
type DayRepository interface {
	FindAll() ([]*entities.DayInfo, error)
	FindByNumber(number int) (*entities.DayInfo, error)
	Save(dayInfo *entities.DayInfo) error
}

// DayLoader defines the interface for loading day implementations
type DayLoader interface {
	Load(dayNumber int) (entities.Day, error)
	DiscoverDays() ([]int, error)
}

// FileReader defines the interface for reading files
type FileReader interface {
	ReadLines(path string) ([]string, error)
}
