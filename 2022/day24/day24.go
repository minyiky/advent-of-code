package day24

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

var cardinalDirs = []aocutils.Vector{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Blizzard struct {
	Pos      aocutils.Vector
	StartPos aocutils.Vector
	Dir      aocutils.Vector
}

type BlizzardFactory struct {
	width, height int
}

func (bf BlizzardFactory) New(pos aocutils.Vector, dirChar rune) Blizzard {
	var dir, startPos aocutils.Vector
	switch dirChar {
	case '>':
		dir = aocutils.NewVector(1, 0)
		startPos = aocutils.NewVector(1, pos.Y)
	case '^':
		dir = aocutils.NewVector(0, 1)
		startPos = aocutils.NewVector(pos.X, 2-bf.height)
	case '<':
		dir = aocutils.NewVector(-1, 0)
		startPos = aocutils.NewVector(bf.width-2, pos.Y)
	case 'v':
		dir = aocutils.NewVector(0, -1)
		startPos = aocutils.NewVector(pos.X, -1)
	}
	return Blizzard{
		Pos:      pos,
		StartPos: startPos,
		Dir:      dir,
	}
}

func moveBlizzards(blizzards []Blizzard, boundry map[aocutils.Vector]bool) map[aocutils.Vector]bool {
	blizMap := make(map[aocutils.Vector]bool)
	for i := range blizzards {
		b := blizzards[i]
		newPos := b.Pos.Add(b.Dir)
		if boundry[newPos] {
			newPos = b.StartPos
		}
		blizzards[i].Pos = newPos
		blizMap[newPos] = true
	}
	return blizMap
}

func makeMove(moveMap, blizMap, boundry map[aocutils.Vector]bool, end aocutils.Vector) (map[aocutils.Vector]bool, bool) {
	newMoves := make(map[aocutils.Vector]bool)
	for pos, _ := range moveMap {
		if !blizMap[pos] {
			newMoves[pos] = true
		}
		for _, dir := range cardinalDirs {
			newPos := pos.Add(dir)
			if newPos == end {
				return nil, true
			}

			if !boundry[newPos] && !blizMap[newPos] {
				newMoves[newPos] = true
			}
		}
	}
	return newMoves, false
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 24 --\n")
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
