package day02

import (
	"log"
	"strings"
)

var playDiff = map[string]int{
	"X": 2, "Y": 0, "Z": 1,
}

var numToPlay = map[int]string{
	1: "X", 2: "Y", 3: "Z",
}

func calculateplay(plays []string) int {
	diff := playDiff[plays[1]]

	play := (playPoints[plays[0]] + diff) % 3
	if play == 0 {
		play = 3
	}

	return calculatePoints([]string{plays[0], numToPlay[play]})
}

func Part2Val(lines []string) (int, error) {
	var score int

	for _, line := range lines {
		plays := strings.Split(line, " ")
		score += calculateplay(plays)
	}

	return score, nil
}

func Part2(lines []string) error {
	score, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("Now that you know what the strategy is meant to be, you would score a total of %d points\n", score)
	return nil
}
