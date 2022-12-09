package day09

import (
	"log"
)

func Part2Val(lines []string) (int, error) {
	return simulateKnots(lines, 10)
}

func Part2(lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("With the bridge broken and the end of of rope whipping around, the end of the rope covered %d squares in the grid", value)
	return nil
}
