package day14

import (
	"fmt"
	"io"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func sandFall2(start aocutils.Vector, blocked map[aocutils.Vector]bool, yLim int) bool {
	sand := start
	for {
		if sand.Y == yLim {
			blocked[sand] = true
			break
		}

		if _, ok := blocked[aocutils.NewVector(sand.X, sand.Y+1)]; !ok {
			sand = aocutils.NewVector(sand.X, sand.Y+1)
			continue
		}

		if _, ok := blocked[aocutils.NewVector(sand.X-1, sand.Y+1)]; !ok {
			sand = aocutils.NewVector(sand.X-1, sand.Y+1)
			continue
		}

		if _, ok := blocked[aocutils.NewVector(sand.X+1, sand.Y+1)]; !ok {
			sand = aocutils.NewVector(sand.X+1, sand.Y+1)
			continue
		}

		if sand == start {
			return true
		}

		blocked[sand] = true
		break
	}

	return false
}

func Part2Val(lines []string) (int, error) {
	var value int

	blocked, yLim, err := createMap(lines)
	if err != nil {
		return 0, err
	}
	yLim += 1

	sandStart := aocutils.NewVector(500, 0)
	for {
		value++
		if out := sandFall2(sandStart, blocked, yLim); out {
			break
		}
	}

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Now that you have seen there is a floor, you realiase the hole that the sand is pooring through will be covered after %d grains of sand have fallen\n", value)
	return nil
}
