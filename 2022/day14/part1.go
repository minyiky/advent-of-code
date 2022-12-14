package day14

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func sandFall(start aocutils.Vector, blocked map[aocutils.Vector]bool, yLim int) bool {
	sand := start
	for {
		if sand.Y == yLim {
			return true
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

		blocked[sand] = true
		break
	}

	return false
}

func Part1Val(lines []string) (int, error) {
	var value int

	blocked, yLim, err := createMap(lines)
	if err != nil {
		return 0, err
	}

	sandStart := aocutils.NewVector(500, 0)
	for {
		if out := sandFall(sandStart, blocked, yLim); out {
			break
		}
		value++
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "While analysing the paths through the cave you estimate that %d grains of sand will fall\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
