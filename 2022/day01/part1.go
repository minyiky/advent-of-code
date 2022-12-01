package day01

import (
	"log"
	"strconv"
)

func Part1(lines []string) error {
	var elfTotal, elfMax int

	for _, line := range lines {
		if line == "" {
			if elfTotal > elfMax {
				elfMax = elfTotal
			}
			elfTotal = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		elfTotal += val
	}

	log.Printf("The elf carrying the most food had %d calories", elfMax)
	return nil
}
