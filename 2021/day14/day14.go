package day14

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func run(lines []string, num int) (int, error) {
	var pairMap = make(map[string]int)

	var instructionMap = make(map[string]rune)

	var lastChar = lines[0][len(lines[0])-1:]

	for i := range lines[0][1:] {
		subStr := lines[0][i : i+2]
		if _, ok := pairMap[subStr]; !ok {
			pairMap[subStr] = 0
		}
		pairMap[subStr] += 1
	}

	for _, line := range lines[2:] {
		var instruction string
		var insert rune
		fmt.Sscanf(line, "%s -> %c", &instruction, &insert)
		instructionMap[instruction] = insert
	}

	for i := 0; i < num; i++ {
		newPairMap := make(map[string]int)

		for pair, number := range pairMap {
			insert := instructionMap[pair]

			subStr := pair[0:1] + string(insert)
			if _, ok := newPairMap[subStr]; !ok {
				newPairMap[subStr] = 0
			}
			newPairMap[subStr] += number

			subStr = string(insert) + pair[1:2]
			if _, ok := newPairMap[subStr]; !ok {
				newPairMap[subStr] = 0
			}
			newPairMap[subStr] += number
		}

		pairMap = newPairMap
	}

	valueMap := make(map[string]int)

	valueMap[lastChar] = 1

	for pair, indices := range pairMap {
		char := pair[0:1]
		if _, ok := valueMap[char]; !ok {
			valueMap[char] = 0
		}
		valueMap[char] += indices
	}

	values := make([]int, 0, len(valueMap))

	for _, indices := range valueMap {
		values = append(values, indices)
	}

	slices.Sort(values)

	return values[len(values)-1] - values[0], nil
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2021 day 14 --\n")
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
