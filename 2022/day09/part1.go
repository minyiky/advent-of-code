package day09

import (
	"log"
)

func Part1Val(lines []string) (int, error) {
	return simulateKnots(lines, 2)
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The knotted tail of the rope covered %d of the squares in the gird", value)
	return nil
}
