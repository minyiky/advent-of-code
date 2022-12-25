package day23

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

//go:embed input.txt
var input string

var NoMove = aocutils.Vector{0, 0}

type Moves struct {
	Directions [][]aocutils.Vector
}

func (m *Moves) NextCycle() {
	m.Directions = append(m.Directions[1:], m.Directions[0])
}

type Elf struct {
	Pos  aocutils.Vector
	Move aocutils.Vector
}

func GetElves(lines []string) []Elf {
	elves := make([]Elf, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				elves = append(elves, Elf{
					Pos:  aocutils.NewVector(x, -y),
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
