package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

// JSONDayRepository stores day info in a JSON file
type JSONDayRepository struct {
	filePath string
	data     map[int]*entities.DayInfo
}

// NewJSONDayRepository creates a new JSON-based day repository
func NewJSONDayRepository(filePath string) (*JSONDayRepository, error) {
	repo := &JSONDayRepository{
		filePath: filePath,
		data:     make(map[int]*entities.DayInfo),
	}

	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Load existing data if file exists
	if err := repo.load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load existing data: %w", err)
	}

	return repo, nil
}

// FindAll returns all day info records
func (r *JSONDayRepository) FindAll() ([]*entities.DayInfo, error) {
	result := make([]*entities.DayInfo, 0, len(r.data))
	for _, dayInfo := range r.data {
		result = append(result, dayInfo)
	}
	return result, nil
}

// FindByNumber finds day info by day number
func (r *JSONDayRepository) FindByNumber(number int) (*entities.DayInfo, error) {
	dayInfo, exists := r.data[number]
	if !exists {
		return nil, fmt.Errorf("day %d not found", number)
	}
	return dayInfo, nil
}

// Save saves or updates day info
func (r *JSONDayRepository) Save(dayInfo *entities.DayInfo) error {
	r.data[dayInfo.Number] = dayInfo
	return r.persist()
}

// load reads data from the JSON file
func (r *JSONDayRepository) load() error {
	file, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&r.data)
}

// persist writes data to the JSON file
func (r *JSONDayRepository) persist() error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(r.data); err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	return nil
}
