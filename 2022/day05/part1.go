package day05

import (
	"log"
)

func Part1Val(lines []string) (string, error) {
	stacks, lines := findStacks(lines)
	var item string
	for _, line := range lines {
		number, start, end := parseInstructions(line)
		for j := 0; j < number; j++ {
			stacks[start], item = stacks[start][:len(stacks[start])-1], stacks[start][len(stacks[start])-1]
			stacks[end] = append(stacks[end], item)
		}
	}
	return topItems(stacks), nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The items on the top of the stacks will be: %s", value)
	return nil
}
