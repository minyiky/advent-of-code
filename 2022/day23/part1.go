package day23

import (
	"fmt"
	"io"
	"math"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func MapElves(elves []Elf) map[point.Point2D]bool {
	grid := make(map[point.Point2D]bool)
	for _, elf := range elves {
		grid[elf.Pos] = true
	}
	return grid
}

func MoveElves(elves []Elf, moves Moves) bool {
	grid := MapElves(elves)

	var moved bool

	plannedMoves := make(map[point.Point2D]int)

	for i, elf := range elves {
		posMoves := make([]point.Point2D, 0, 4)
	dir:
		for _, dir := range moves.Directions {
			for _, subDir := range dir {
				if grid[point.Add(elf.Pos, subDir)] {
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
		plannedMoves[point.Add(elf.Pos, posMoves[0])]++
	}

	for elf := range elves {
		if plannedMoves[point.Add(elves[elf].Pos, elves[elf].Move)] == 1 {
			elves[elf].Pos = point.Add(elves[elf].Pos, elves[elf].Move)
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
		if elf.Pos.X() < xMin {
			xMin = elf.Pos.X()
		}
		if elf.Pos.Y() < yMin {
			yMin = elf.Pos.Y()
		}
		if elf.Pos.X() > xMax {
			xMax = elf.Pos.X()
		}
		if elf.Pos.Y() > yMax {
			yMax = elf.Pos.Y()
		}
	}
	return xMin, xMax + 1, yMin, yMax + 1
}

func Part1Val(lines []string) (int, error) {
	var value int

	moves := Moves{
		[][]point.Point2D{
			{point.NewPoint2D(0, 1), point.NewPoint2D(-1, 1), point.NewPoint2D(1, 1)},
			{point.NewPoint2D(0, -1), point.NewPoint2D(-1, -1), point.NewPoint2D(1, -1)},
			{point.NewPoint2D(-1, 0), point.NewPoint2D(-1, -1), point.NewPoint2D(-1, 1)},
			{point.NewPoint2D(1, 0), point.NewPoint2D(1, -1), point.NewPoint2D(1, 1)},
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
