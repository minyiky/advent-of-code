package day06

import (
	"log"
)

func Part1Val(line string) (int, error) {
	return findSequence(line, 4)
}

func Part1(line string) error {
	value, err := Part1Val(line)
	if err != nil {
		return err
	}
	log.Printf("A marker sequence of 4 characters was found after receiving %d characters", value)
	return nil
}
