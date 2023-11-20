package day17

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

type Shape struct {
	name      string
	points    []point.Point2D
	highPoint int
}

type ShapeList struct {
	Shapes []Shape
	index  int
}

func (s *ShapeList) Next() Shape {
	shape := Shape{
		name:      s.Shapes[s.index].name,
		points:    container.CopySlice(s.Shapes[s.index].points),
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
		points: []point.Point2D{
			point.NewPoint2D(0, 0),
			point.NewPoint2D(1, 0),
			point.NewPoint2D(2, 0),
			point.NewPoint2D(3, 0),
		},
		highPoint: 0,
	}
	Cross = Shape{
		name: "cross",
		points: []point.Point2D{
			point.NewPoint2D(1, 0),
			point.NewPoint2D(0, 1),
			point.NewPoint2D(2, 1),
			point.NewPoint2D(1, 2),
		},
		highPoint: 3,
	}
	Tall = Shape{
		name: "tall",
		points: []point.Point2D{
			point.NewPoint2D(0, 0),
			point.NewPoint2D(0, 1),
			point.NewPoint2D(0, 2),
			point.NewPoint2D(0, 3),
		},
		highPoint: 3,
	}
	L = Shape{
		name: "l",
		points: []point.Point2D{
			point.NewPoint2D(0, 0),
			point.NewPoint2D(1, 0),
			point.NewPoint2D(2, 0),
			point.NewPoint2D(2, 1),
			point.NewPoint2D(2, 2),
		},
		highPoint: 4,
	}
	Box = Shape{
		name: "box",
		points: []point.Point2D{
			point.NewPoint2D(0, 0),
			point.NewPoint2D(1, 0),
			point.NewPoint2D(0, 1),
			point.NewPoint2D(1, 1),
		},
		highPoint: 3,
	}
)

func QuickHeight(num int, line string) int {
	if num <= 1741 {
		grid := make(map[point.Point2D]bool)
		for i := 0; i < 7; i++ {
			grid[point.NewPoint2D(i, -1)] = true
		}
		return HeightAfterFall(num, line, grid, point.NewPoint2D(0, 0), 0, 0)
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

	grid := make(map[point.Point2D]bool)
	grid[point.NewPoint2D(0, height-3)] = true
	grid[point.NewPoint2D(1, height-2)] = true
	grid[point.NewPoint2D(2, height-1)] = true
	grid[point.NewPoint2D(3, height-1)] = true
	grid[point.NewPoint2D(4, height-1)] = true
	grid[point.NewPoint2D(4, height-2)] = true
	grid[point.NewPoint2D(4, height-3)] = true
	grid[point.NewPoint2D(4, height-4)] = true
	grid[point.NewPoint2D(4, height-5)] = true
	grid[point.NewPoint2D(5, height-1)] = true
	grid[point.NewPoint2D(5, height-6)] = true

	return HeightAfterFall(num-numFallen, line, grid, point.NewPoint2D(-1, 0), height, 1)

}

func HeightAfterFall(num int, line string, grid map[point.Point2D]bool, initialPush point.Point2D, startHeight, startIndex int) int {
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
				shape.points[j] = point.Add(pos, initialPush)
			}
		}
		for j, pos := range shape.points {
			shape.points[j] = point.Add(pos, point.NewPoint2D(2, height))
		}
		for {
			charIndex += 1
			charIndex %= lineLen
			char := line[charIndex]
			push := point.NewPoint2D(1, 0)
			if char == '<' {
				push = point.NewPoint2D(-1, 0)
			}
			if !blockedSide(shape, push, grid) {
				for j, pos := range shape.points {
					shape.points[j] = point.Add(pos, push)
				}
			}

			if height <= value && blockedDown(shape, grid) {
				for _, pos := range shape.points {
					grid[pos] = true
				}
				blockHeight := shape.points[shape.highPoint].Y() + 1
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
				shape.points[j] = point.Add(pos, point.NewPoint2D(0, -1))
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
