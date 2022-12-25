package day24

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func FindExit2(start, end aocutils.Vector, blizzards []Blizzard, boundry map[aocutils.Vector]bool) int {
	var moves int
	moveMap := map[aocutils.Vector]bool{start: true}

	for {
		moves++
		blizMap := moveBlizzards(blizzards, boundry)

		newMoves, ok := makeMove(moveMap, blizMap, boundry, end)
		if ok {
			break
		}

		moveMap = newMoves
	}

	moveMap = map[aocutils.Vector]bool{end: true}

	for {
		moves++
		blizMap := moveBlizzards(blizzards, boundry)

		newMoves, ok := makeMove(moveMap, blizMap, boundry, start)
		if ok {
			break
		}

		moveMap = newMoves
	}

	moveMap = map[aocutils.Vector]bool{start: true}

	for {
		moves++
		blizMap := moveBlizzards(blizzards, boundry)

		newMoves, ok := makeMove(moveMap, blizMap, boundry, end)
		if ok {
			break
		}

		moveMap = newMoves
	}

	return moves
}

func Part2Val(lines []string) (int, error) {
	width := len(lines[0])
	height := len(lines)
	blizzardFactory := BlizzardFactory{width: width, height: height}

	boundry := map[aocutils.Vector]bool{{1, 1}: true}
	blizzards := []Blizzard{}

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '>', '<', '^', 'v':
				blizzards = append(blizzards, blizzardFactory.New(aocutils.NewVector(x, -y), char))
			case '#':
				boundry[aocutils.NewVector(x, -y)] = true
			}
		}
	}

	return FindExit2(
		aocutils.NewVector(1, 0),
		aocutils.NewVector(width-2, 1-height),
		blizzards,
		boundry,
	), nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Accounting for the elve's inepditude, you now reckon that it will take you %d minutes to reach the extraction point\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
