package day17

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

type Shape struct {
	name      string
	points    []aocutils.Vector
	highPoint int
}

type ShapeList struct {
	Shapes []Shape
	index  int
}

func (s *ShapeList) Next() Shape {
	shape := Shape{
		name:      s.Shapes[s.index].name,
		points:    aocutils.CopySlice(s.Shapes[s.index].points),
		highPoint: s.Shapes[s.index].highPoint,
	}
	s.index += 1
	if s.index == len(s.Shapes) {
		s.index = 0
	}
	return shape
}

var (
	Flat = Shape{
		name: "flat",
		points: []aocutils.Vector{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
		},
		highPoint: 0,
	}
	Cross = Shape{
		name: "cross",
		points: []aocutils.Vector{
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 2, Y: 1},
			{X: 1, Y: 2},
		},
		highPoint: 3,
	}
	Tall = Shape{
		name: "tall",
		points: []aocutils.Vector{
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 0, Y: 3},
		},
		highPoint: 3,
	}
	L = Shape{
		name: "l",
		points: []aocutils.Vector{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 1},
			{X: 2, Y: 2},
		},
		highPoint: 4,
	}
	Box = Shape{
		name: "box",
		points: []aocutils.Vector{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
		},
		highPoint: 3,
	}
)

func QuickHeight(num int, line string) int {
	if num <= 1741 {
		grid := make(map[aocutils.Vector]bool)
		for i := 0; i < 7; i++ {
			grid[aocutils.Vector{i, -1}] = true
		}
		return HeightAfterFall(num, line, grid, aocutils.Vector{}, 0, 0)
	}

	initHeight := 2701
	regHeight := 2695
	initFallen := 1741
	regFallen := 1735

	numFallen := initFallen
	height := initHeight
	for i := initFallen + regFallen; i < num; i += regFallen {
		numFallen += regFallen
		height += regHeight
	}

	grid := make(map[aocutils.Vector]bool)
	grid[aocutils.Vector{0, height - 3}] = true
	grid[aocutils.Vector{1, height - 2}] = true
	grid[aocutils.Vector{2, height - 1}] = true
	grid[aocutils.Vector{3, height - 1}] = true
	grid[aocutils.Vector{4, height - 1}] = true
	grid[aocutils.Vector{4, height - 2}] = true
	grid[aocutils.Vector{4, height - 3}] = true
	grid[aocutils.Vector{4, height - 4}] = true
	grid[aocutils.Vector{4, height - 5}] = true
	grid[aocutils.Vector{5, height - 1}] = true
	grid[aocutils.Vector{5, height - 6}] = true

	return HeightAfterFall(num-numFallen, line, grid, aocutils.NewVector(-1, 0), height, 1)

}

func HeightAfterFall(num int, line string, grid map[aocutils.Vector]bool, initialPush aocutils.Vector, startHeight, startIndex int) int {
	value := startHeight
	shapes := ShapeList{
		[]Shape{Flat, Cross, L, Tall, Box},
		startIndex,
	}

	charIndex := -1
	lineLen := len(line)

	for i := 0; i < num; i++ {
		var height int
		if i == 0 && num < 1741 {
			height = value + 2
		} else {
			height = value + 3
		}
		shape := shapes.Next()
		if i == 0 {
			for j, pos := range shape.points {
				shape.points[j] = pos.Add(initialPush)
			}
		}
		for j, pos := range shape.points {
			shape.points[j] = pos.Add(aocutils.NewVector(2, height))
		}
		for {
			charIndex += 1
			charIndex %= lineLen
			char := line[charIndex]
			push := aocutils.NewVector(1, 0)
			if char == '<' {
				push = aocutils.NewVector(-1, 0)
			}
			if !blockedSide(shape, push, grid) {
				for j, pos := range shape.points {
					shape.points[j] = pos.Add(push)
				}
			}

			if height <= value && blockedDown(shape, grid) {
				for _, pos := range shape.points {
					grid[pos] = true
				}
				blockHeight := shape.points[shape.highPoint].Y + 1
				if blockHeight > value {
					value = blockHeight
				}
				break
			}
			if height < (value - 200) {
				fmt.Println(shape.name)
				break
			}
			for j, pos := range shape.points {
				shape.points[j] = pos.Add(aocutils.NewVector(0, -1))
			}
			height -= 1
		}

	}

	return value
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")

	fmt.Fprintf(w, "-- Solution for 2022 day 17 --\n")
	if err := Part1(w, input); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, input); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
