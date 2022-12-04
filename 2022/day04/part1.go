package day04

import (
	"errors"
	"log"
	"strings"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		elfs := strings.Split(line, ",")
		if len(elfs) != 2 {
			return 0, errors.New("wrong number of sections found")
		}

		elf1, err := getRange(elfs[0])
		if len(elfs) != 2 {
			return 0, err
		}

		elf2, err := getRange(elfs[1])
		if len(elfs) != 2 {
			return 0, err
		}

		if (elf1[0] <= elf2[0] && elf1[1] >= elf2[1]) ||
			(elf1[0] >= elf2[0] && elf1[1] <= elf2[1]) {
			value += 1
			continue
		}

	}

	return value, nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("For %d of the pairs of elves, ones tasks were completely included in the other", value)
	return nil
}
