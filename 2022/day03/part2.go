package day03

import (
	"log"
)

func getMidMatches(bag string, items map[rune]bool) map[rune]bool {
	var midMatches = make(map[rune]bool)

	for _, item := range bag {
		if _, ok := items[item]; ok {
			midMatches[item] = true
		}
	}

	return midMatches
}

func Part2Val(lines []string) (int, error) {
	var score int

	for i := 0; i < len(lines); i += 3 {
		itemsFirst := getItems(lines[i])
		itemsMid := getMidMatches(lines[i+1], itemsFirst)
		badge, err := getMatch(lines[i+2], itemsMid)
		if err != nil {
			return 0, err
		}

		val, err := getPriority(badge)
		if err != nil {
			return 0, err
		}

		score += val

	}

	return score, nil
}

func Part2(lines []string) error {
	score, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The badges marking the different groups had a total priority of %d points", score)
	return nil
}
