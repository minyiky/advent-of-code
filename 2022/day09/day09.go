package day09

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

type Vector struct {
	x int
	y int
}

var dirMap = map[rune]Vector{
	'U': {0, 1},
	'D': {0, -1},
	'R': {1, 0},
	'L': {-1, 0},
}

func (v *Vector) move(dir Vector) {
	v.x += dir.x
	v.y += dir.y
}

func (v *Vector) follow(leader Vector) {
	xDist, yDist := aocutils.Abs(leader.x-v.x), aocutils.Abs(leader.y-v.y)
	if xDist <= 1 && yDist <= 1 {
		return
	}

	if leader.x > v.x {
		v.x++
	}
	if leader.x < v.x {
		v.x--
	}
	if leader.y > v.y {
		v.y++
	}
	if leader.y < v.y {
		v.y--
	}
}

func simulateKnots(lines []string, num int) (int, error) {
	var value int

	knots := make([]Vector, num)
	tailPos := make(map[Vector]bool)

	dirRune, count := ' ', 0
	for _, line := range lines {
		if _, err := fmt.Sscanf(line, "%c %d", &dirRune, &count); err != nil {
			return 0, err
		}

		dir := dirMap[dirRune]
		for i := 0; i < count; i++ {
			// Move the head knot
			knots[0].move(dir)

			// Make all of the other knots follow
			for k := 1; k < num; k++ {
				knots[k].follow(knots[k-1])
			}
			if _, ok := tailPos[knots[num-1]]; !ok {
				tailPos[knots[num-1]] = true
				value++
			}
		}
	}

	return value, nil
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 09 --\n")
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
