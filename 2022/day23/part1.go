package day23

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func MapElves(elves []Elf) map[aocutils.Vector]bool {
	grid := make(map[aocutils.Vector]bool)
	for _, elf := range elves {
		grid[elf.Pos] = true
	}
	return grid
}

func MoveElves(elves []Elf, moves Moves) bool {
	grid := MapElves(elves)

	var moved bool

	plannedMoves := make(map[aocutils.Vector]int)

	for i, elf := range elves {
		posMoves := make([]aocutils.Vector, 0, 4)
	dir:
		for _, dir := range moves.Directions {
			for _, subDir := range dir {
				if grid[elf.Pos.Add(subDir)] {
					continue dir
				}
			}
			posMoves = append(posMoves, dir[0])
		}
		if len(posMoves) == 4 || len(posMoves) == 0 {
			continue
		}
		moved = true
		elves[i].Move = posMoves[0]
		plannedMoves[elf.Pos.Add(posMoves[0])]++
	}

	for elf := range elves {
		if plannedMoves[elves[elf].Pos.Add(elves[elf].Move)] == 1 {
			elves[elf].Pos = elves[elf].Pos.Add(elves[elf].Move)
		}
		elves[elf].Move = NoMove
	}

	return moved
}

func findBounds(elves []Elf) (int, int, int, int) {
	xMax := math.MinInt
	xMin := math.MaxInt
	yMax := math.MinInt
	yMin := math.MaxInt

	for _, elf := range elves {
		if elf.Pos.X < xMin {
			xMin = elf.Pos.X
		}
		if elf.Pos.Y < yMin {
			yMin = elf.Pos.Y
		}
		if elf.Pos.X > xMax {
			xMax = elf.Pos.X
		}
		if elf.Pos.Y > yMax {
			yMax = elf.Pos.Y
		}
	}
	return xMin, xMax + 1, yMin, yMax + 1
}

func Part1Val(lines []string) (int, error) {
	var value int

	moves := Moves{
		[][]aocutils.Vector{
			{{0, 1}, {-1, 1}, {1, 1}},
			{{0, -1}, {-1, -1}, {1, -1}},
			{{-1, 0}, {-1, -1}, {-1, 1}},
			{{1, 0}, {1, -1}, {1, 1}},
		},
	}
	elves := GetElves(lines)
	for i := 0; i < 10; i++ {
		// fmt.Printf("== End of Round %d ==\n", i)
		MoveElves(elves, moves)
		moves.NextCycle()
	}
	xMin, xMax, yMin, yMax := findBounds(elves)
	fmt.Printf("xMin: %d, xMax: %d, yMin: %d, yMax: %d\n\n\n", xMin, xMax, yMin, yMax)

	value = ((xMax - xMin) * (yMax - yMin)) - len(elves)
	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "There are %d free squares in the area covered by the elves\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
