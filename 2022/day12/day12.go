package day12

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

//go:embed input.txt
var input string

func findSummit(position, end aocutils.Vector, steps int, grid [][]rune, vistited map[aocutils.Vector]int) (int, bool) {
	if position == end {
		return steps, true
	}

	vistited[position] = steps

	steps++

	char := grid[position.Y][position.X]
	up, down, right, left := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	var upOk, downOk, rightOk, leftOk, ok, seen bool
	var last int

	// Look Up
	last, seen = vistited[aocutils.NewVector(position.X, position.Y-1)]
	ok = (!seen || last > steps)
	if position.Y != 0 &&
		checkHeight(char, grid[position.Y-1][position.X]) &&
		ok {
		up, upOk = findSummit(aocutils.NewVector(position.X, position.Y-1), end, steps, grid, vistited)
	}
	// Look Down
	last, seen = vistited[aocutils.NewVector(position.X, position.Y+1)]
	ok = (!seen || last > steps)
	if position.Y != len(grid)-1 &&
		checkHeight(char, grid[position.Y+1][position.X]) &&
		ok {
		down, downOk = findSummit(aocutils.NewVector(position.X, position.Y+1), end, steps, grid, vistited)
	}
	// Look Left
	last, seen = vistited[aocutils.NewVector(position.X-1, position.Y)]
	ok = (!seen || last > steps)
	if position.X != 0 &&
		checkHeight(char, grid[position.Y][position.X-1]) &&
		ok {
		left, leftOk = findSummit(aocutils.NewVector(position.X-1, position.Y), end, steps, grid, vistited)
	}
	last, seen = vistited[aocutils.NewVector(position.X+1, position.Y)]
	ok = (!seen || last > steps)
	// Look Right
	if position.X != len(grid[0])-1 &&
		checkHeight(char, grid[position.Y][position.X+1]) &&
		ok {
		right, rightOk = findSummit(aocutils.NewVector(position.X+1, position.Y), end, steps, grid, vistited)
	}

	if !upOk && !downOk && !leftOk && !rightOk {
		return math.MaxInt, false
	}
	heights := sort.IntSlice{up, down, left, right}
	sort.Sort(heights)
	return heights[0], true
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 12 --\n")
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
