package day05

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

//go:embed input.txt
var input string

func findStacks(lines []string) ([][]string, []string, error) {
	if (len(lines[0])+1)%4 != 0 {
		return nil, nil, errors.New("incorrect fomat for stacks")
	}

	var numStacks int = (len(lines[0]) + 1) / 4
	stacks := make([][]string, numStacks)
	for i, line := range lines {
		// This indicates that we are at the Stack number row
		if string(line[1]) == "1" {
			// if the instructions dont exist then send back an empty list for lines
			if len(lines) < 1+2 {
				return stacks, []string{}, nil
			}

			// Reset the lines slice to only include instructions and break the loop
			lines = lines[i+2:]
			break
		}

		for j := 0; j < numStacks; j++ {
			// If the character is not a space insert a
			if char := string(line[4*j+1]); char != " " {
				stacks[j] = append(stacks[j], char)
			}
		}
	}

	// Reverse the stacks so the first items are at the top
	for _, stack := range stacks {
		aocutils.ReverseSlice(stack)
	}

	return stacks, lines, nil
}

func parseInstructions(line string) (int, int, int, error) {
	var number, start, end int
	if _, err := fmt.Sscanf(line, "move %d from %d to %d", &number, &start, &end); err != nil {
		return 0, 0, 0, err
	}
	return number, start - 1, end - 1, nil
}

func topItems(stacks [][]string) string {
	var items string
	for j := 0; j < len(stacks); j++ {
		items += stacks[j][len(stacks[j])-1]
	}
	return items
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 05 --\n")
	if err := Part1(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
