package day04

import (
	"log"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines{
	}

	return value, nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The value found was: %d", value)
	return nil
}
