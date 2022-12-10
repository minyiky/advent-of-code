package day05

import (
	"fmt"
	"io"
)

func Part2Val(lines []string) (string, error) {
	stacks, lines, err := findStacks(lines)
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		var item []string
		number, start, end, err := parseInstructions(line)
		if err != nil {
			return "", err
		}

		if number > len(stacks[start]) {
			return "", fmt.Errorf("trying to move %d crates in a stack of size %d", number, len(stacks[start]))
		}

		stacks[start], item = stacks[start][:len(stacks[start])-number], stacks[start][len(stacks[start])-number:]
		stacks[end] = append(stacks[end], item...)
	}
	return topItems(stacks), nil
}

func Part2(w io.Writer, lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Now that the correct version of the crane is know, the items on the top of the stacks will be: %s\n", value)
	return nil
}
