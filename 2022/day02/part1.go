package day02

import (
	"log"
	"strings"
)

func calculatePoints(plays []string) int {
	points := playPoints[plays[1]]

	if playPoints[plays[1]] == playPoints[plays[0]] {
		return points + 3
	}

	if playPoints[plays[1]] == playPoints[plays[0]]+1 || playPoints[plays[1]] == playPoints[plays[0]]-2 {
		return points + 6
	}
	return points
}

func Part1(lines []string) error {
	var score int

	for _, line := range lines {
		plays := strings.Split(line, " ")
		score += calculatePoints(plays)
	}

	log.Printf("Following the guide the elf gave you, you would score a total of %d points", score)
	return nil
}
