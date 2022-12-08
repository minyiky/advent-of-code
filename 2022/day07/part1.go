package day07

import (
	"log"
)

func Part1Val(lines []string) (int, error) {
	var value int

	dirs := ReadDirs(lines)

	for _, v := range dirs {
		if v < 100000 {
			value += v
		}
	}

	return value, nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The total soze of directories under 100000 was %d", value)
	return nil
}
