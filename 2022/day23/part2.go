package day23

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func Part2Val(lines []string) (int, error) {
	moves := Moves{
		[][]aocutils.Vector{
			{{0, 1}, {-1, 1}, {1, 1}},
			{{0, -1}, {-1, -1}, {1, -1}},
			{{-1, 0}, {-1, -1}, {-1, 1}},
			{{1, 0}, {1, -1}, {1, 1}},
		},
	}
	elves := GetElves(lines)
	var round int
	for {
		// fmt.Printf("== End of Round %d ==\n", i)
		round++
		if moved := MoveElves(elves, moves); !moved {
			break
		}
		moves.NextCycle()
	}

	return round, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "It will take %d minutes to reach the final position\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
