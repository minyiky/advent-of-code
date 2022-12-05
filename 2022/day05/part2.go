package day05

import (
	"log"
)

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines{
	}

	return value, nil
}

func Part2(lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The value found was: %d", value)
	return nil
}
