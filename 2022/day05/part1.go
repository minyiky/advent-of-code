package day05

import (
	"fmt"
	"log"
)

func Part1Val(lines []string) (string, error) {
	stacks, lines, err := findStacks(lines)
	if err != nil {
		return "", err
	}
	var item string
	for _, line := range lines {
		number, start, end, err := parseInstructions(line)
		if err != nil {
			return "", err
		}

		if number > len(stacks[start]) {
			return "", fmt.Errorf("trying to move %d crates in a stack of size %d", number, len(stacks[start]))
		}

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
