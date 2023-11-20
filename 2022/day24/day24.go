package day24

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

var cardinalDirs = []point.Point2D{
	point.NewPoint2D(1, 0),
	point.NewPoint2D(-1, 0),
	point.NewPoint2D(0, 1),
	point.NewPoint2D(0, -1),
}

type Blizzard struct {
	Pos      point.Point2D
	StartPos point.Point2D
	Dir      point.Point2D
}

type BlizzardFactory struct {
	width, height int
}

func (bf BlizzardFactory) New(pos point.Point2D, dirChar rune) Blizzard {
	var dir, startPos point.Point2D
	switch dirChar {
	case '>':
		dir = point.NewPoint2D(1, 0)
		startPos = point.NewPoint2D(1, pos.Y())
	case '^':
		dir = point.NewPoint2D(0, 1)
		startPos = point.NewPoint2D(pos.X(), 2-bf.height)
	case '<':
		dir = point.NewPoint2D(-1, 0)
		startPos = point.NewPoint2D(bf.width-2, pos.Y())
	case 'v':
		dir = point.NewPoint2D(0, -1)
		startPos = point.NewPoint2D(pos.X(), -1)
	}
	return Blizzard{
		Pos:      pos,
		StartPos: startPos,
		Dir:      dir,
	}
}

func moveBlizzards(blizzards []Blizzard, boundry map[point.Point2D]bool) map[point.Point2D]bool {
	blizMap := make(map[point.Point2D]bool)
	for i := range blizzards {
		b := blizzards[i]
		newPos := point.Add(b.Pos, b.Dir)
		if boundry[newPos] {
			newPos = b.StartPos
		}
		blizzards[i].Pos = newPos
		blizMap[newPos] = true
	}
	return blizMap
}

func makeMove(moveMap, blizMap, boundry map[point.Point2D]bool, end point.Point2D) (map[point.Point2D]bool, bool) {
	newMoves := make(map[point.Point2D]bool)
	for pos, _ := range moveMap {
		if !blizMap[pos] {
			newMoves[pos] = true
		}
		for _, dir := range cardinalDirs {
			newPos := point.Add(pos, dir)
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
