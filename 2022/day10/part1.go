package day10

import (
	"fmt"
	"log"
	"strings"
)

func Part1Val(lines []string) (int, error) {
	var value int
	var cycle int
	x := 1
	for _, line := range lines {
		cycle++
		if (cycle-20)%40 == 0 {
			value += x * cycle
		}

		if strings.HasPrefix(line, "noop") {
			continue
		}

		cycle++

		if (cycle-20)%40 == 0 {
			value += x * cycle
		}

		var number int
		if _, err := fmt.Sscanf(line, "addx %d", &number); err != nil {
			return 0, err
		}

		x += number
	}

	return value, nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("Analyisinf the register, it looks like the cpu has a total signal strength of %d", value)
	return nil
}
