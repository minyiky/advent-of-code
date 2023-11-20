package day23

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

var NoMove = point.NewPoint2D(0, 0)

type Moves struct {
	Directions [][]point.Point2D
}

func (m *Moves) NextCycle() {
	m.Directions = append(m.Directions[1:], m.Directions[0])
}

type Elf struct {
	Pos  point.Point2D
	Move point.Point2D
}

func GetElves(lines []string) []Elf {
	elves := make([]Elf, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				elves = append(elves, Elf{
					Pos:  point.NewPoint2D(x, -y),
					Move: NoMove,
				})
			}
		}
	}
	return elves
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 23 --\n")
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
