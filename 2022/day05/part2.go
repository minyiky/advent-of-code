package day05

import (
	"log"
)

func Part2Val(lines []string) (string, error) {
	stacks, lines := findStacks(lines)
	for _, line := range lines {
		var item []string
		number, start, end := parseInstructions(line)
		stacks[start], item = stacks[start][:len(stacks[start])-number], stacks[start][len(stacks[start])-number:]
		stacks[end] = append(stacks[end], item...)
	}
	return topItems(stacks), nil
}

func Part2(lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("Now that the correct version of the crane is know, the items on the top of the stacks will be: %s", value)
	return nil
}
