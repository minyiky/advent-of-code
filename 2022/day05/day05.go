package day05

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func findStacks(lines []string) ([][]string, []string) {
	var numStacks int = (len(lines[0]) + 1) / 4
	stacks := make([][]string, numStacks)
	for i, line := range lines {
		if string(line[1]) == "1" {
			lines = lines[i+2:]
			break
		}

		for j := 0; j < numStacks; j++ {
			if char := string(line[4*j+1]); char != " " {
				stacks[j] = append([]string{char}, stacks[j]...)
			}
		}
	}

	return stacks, lines
}

func parseInstructions(line string) (int, int, int) {
	var number, start, end int
	fmt.Sscanf(line, "move %d from %d to %d", &number, &start, &end)
	return number, start - 1, end - 1

}

func topItems(stacks [][]string) string {
	var items string
	for j := 0; j < len(stacks); j++ {
		items += stacks[j][len(stacks[j])-1]
	}
	return items
}

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 05 --")
	if err := Part1(lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
