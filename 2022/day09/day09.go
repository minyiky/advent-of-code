package day09

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

var dirMap = map[rune]point.Point2D{
	'U': point.NewPoint2D(0, 1),
	'D': point.NewPoint2D(0, -1),
	'R': point.NewPoint2D(1, 0),
	'L': point.NewPoint2D(-1, 0),
}

func follow(v, leader point.Point2D) point.Point2D {
	xDist, yDist := maths.Abs(leader.X()-v.X()), maths.Abs(leader.Y()-v.Y())
	if xDist <= 1 && yDist <= 1 {
		return v
	}

	if leader.X() > v.X() {
		v.SetX(v.X() + 1)
	}
	if leader.X() < v.X() {
		v.SetX(v.X() - 1)
	}
	if leader.Y() > v.Y() {

		v.SetY(v.Y() + 1)
	}
	if leader.Y() < v.Y() {
		v.SetY(v.Y() - 1)
	}

	return v
}

func simulateKnots(lines []string, num int) (int, error) {
	var value int

	knots := make([]point.Point2D, num)
	for i, _ := range knots {
		knots[i] = point.NewPoint2D(0, 0)
	}

	tailPos := make(map[point.Point2D]bool)

	dirRune, count := ' ', 0
	for _, line := range lines {
		if _, err := fmt.Sscanf(line, "%c %d", &dirRune, &count); err != nil {
			return 0, err
		}

		dir := dirMap[dirRune]
		for i := 0; i < count; i++ {
			// Move the head knot
			knots[0] = point.Add(knots[0], dir)

			// Make all of the other knots follow
			for k := 1; k < num; k++ {
				knots[k] = follow(knots[k], knots[k-1])
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
