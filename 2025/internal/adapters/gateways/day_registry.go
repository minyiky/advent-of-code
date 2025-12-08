package gateways

import (
	"fmt"
	"sort"

	"github.com/minyiky/advent-of-code/2025/internal/entities"
)

// DayRegistration holds a day number and its implementation
type DayRegistration struct {
	Number int
	Day    entities.VisualizableDay
}

// DayRegistry manages registration and loading of day implementations
type DayRegistry struct {
	registry map[int]entities.VisualizableDay
}

// NewDayRegistry creates a new day registry with the provided day implementations
func NewDayRegistry(registrations ...entities.VisualizableDay) *DayRegistry {
	registry := make(map[int]entities.VisualizableDay, len(registrations))
	for _, reg := range registrations {
		registry[reg.Number()] = reg
	}
	return &DayRegistry{
		registry: registry,
	}
}

// Load loads a day implementation from the registry
func (r *DayRegistry) Load(dayNumber int) (entities.Day, error) {
	day, exists := r.registry[dayNumber]
	if !exists {
		return nil, fmt.Errorf("day %d not registered", dayNumber)
	}
	return day, nil
}

// DiscoverDays returns all registered day numbers in sorted order
func (r *DayRegistry) DiscoverDays() ([]int, error) {
	days := make([]int, 0, len(r.registry))
	for dayNum := range r.registry {
		days = append(days, dayNum)
	}
	sort.Ints(days)
	return days, nil
}
