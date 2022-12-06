package day06

import (
	"log"
)

func Part2Val(line string) (int, error) {
	return findSequence(line, 14)
}

func Part2(line string) error {
	value, err := Part2Val(line)
	if err != nil {
		return err
	}
	log.Printf("A marker sequence of 14 characters was found after receiving %d characters", value)
	return nil
}
