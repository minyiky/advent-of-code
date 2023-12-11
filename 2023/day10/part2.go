package day10

import (
	"fmt"
	"io"
	"time"

	"github.com/fatih/color"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type limits struct {
	minX, maxX, minY, maxY int
}

var (
	r = color.New(color.FgRed)
	g = color.New(color.FgGreen)
)

func convertChar(input rune) rune {
	switch input {
	case '-':
		return '─'
	case '|':
		return '│'
	case 'L':
		return '└'
	case 'F':
		return '┌'
	case 'J':
		return '┘'
	case '7':
		return '┐'
	}
	return 'X'
}

func (l limits) Outside(p point.Point2D) bool {
	return p.X() < l.minX || p.X() > l.maxX || p.Y() < l.minY || p.Y() > l.maxY
}

func WayOut(pos point.Point2D, blocks, seen, outside map[point.Point2D]bool, limits limits) bool {
	if blocks[pos] {
		return false
	}

	if outside[pos] || limits.Outside(pos) {
		return true
	}

	seen[pos] = true

	for _, direction := range directions {
		newPoint := point.Add(pos, direction)
		if !seen[newPoint] {
			if WayOut(newPoint, blocks, seen, outside, limits) {
				return true
			}
		}
	}
	return false
}

func Part2Val(lines []string) (int, error) {
	var value int

	pipes := make(map[point.Point2D]pipe)
	var start point.Point2D

	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				start = point.NewPoint2D(x, -y)
				continue
			}
			if char == '.' {
				continue
			}
			pipes[point.NewPoint2D(x, -y)] = pipeTypes[char]
		}
	}

	pos := start
	var dir point.Point2D

	var FirstPipe point.Point2D
	var LastPipe point.Point2D

	for _, d := range directions {
		pos = point.Add(start, d)
		var err error
		if dir, err = pipes[pos].move(d); err == nil {
			FirstPipe = d
			break
		}
	}

	for pos != start {
		pos = point.Add(pos, dir)
		dir, _ = pipes[pos].move(dir)
		LastPipe = dir
	}

	startRow := ""

	startPipe := func() rune {
		if dir, _ := pipeTypes['|'].move(FirstPipe); dir == LastPipe {
			return '|'
		}
		if dir, _ := pipeTypes['-'].move(FirstPipe); dir == LastPipe {
			return '-'
		}
		if dir, _ := pipeTypes['J'].move(FirstPipe); dir == LastPipe {
			return 'F'
		}
		if dir, _ := pipeTypes['F'].move(FirstPipe); dir == LastPipe {
			return 'J'
		}
		if dir, _ := pipeTypes['L'].move(FirstPipe); dir == LastPipe {
			return 'L'
		}
		return '7'
	}()

	for _, char := range lines[-start.Y()] {
		if char == 'S' {
			char = startPipe
		}
		startRow += string(char)
	}

	lines[-start.Y()] = startRow

	newLines := make([]string, 0, len(lines))

	for _, line := range lines {
		newLine := ""
		secondLine := ""
		for _, char := range line {
			switch char {
			case '|', '.', 'J', '7':
				newLine += string(char) + "*"
			case '-', 'F', 'L':
				newLine += string(char) + "-"
			}
			switch char {
			case '|', 'F', '7':
				secondLine += "|*"
			default:
				secondLine += "**"
			}
		}
		newLines = append(newLines, newLine)
		newLines = append(newLines, secondLine)
	}

	pipes = make(map[point.Point2D]pipe)

	for y, line := range newLines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			pipes[point.NewPoint2D(x, -y)] = pipeTypes[char]
		}
	}

	start = point.NewPoint2D(2*start.X(), 2*start.Y())
	pos = point.Add(start, FirstPipe)

	dir, _ = pipes[pos].move(FirstPipe)

	loop := map[point.Point2D]bool{
		start: true,
		pos:   true,
	}

	for pos != start {
		pos = point.Add(pos, dir)
		loop[pos] = true
		dir, _ = pipes[pos].move(dir)
		LastPipe = dir
	}

	outside := make(map[point.Point2D]bool)
	inside := make(map[point.Point2D]bool)
	limits := limits{
		minX: 0,
		maxX: 2*len(lines[0]) - 1,
		minY: 2*-len(lines) + 1,
		maxY: 0,
	}

	for y, line := range newLines {
		for x := range line {
			p := point.NewPoint2D(x, -y)
			if _, ok := loop[p]; ok {
				continue
			}
			if inside[p] || outside[p] {
				continue
			}
			visited := make(map[point.Point2D]bool)
			if !WayOut(p, loop, visited, outside, limits) {
				for p := range visited {
					inside[p] = true
				}
			} else {
				for p := range visited {
					outside[p] = true
				}
			}
		}
	}
	for p := range inside {
		if newLines[-p.Y()][p.X()] != '*' {
			if p.X()%2 == 0 && p.Y()%2 == 0 {
				value++
			}
		}
	}

	for y, line := range newLines {
		if y%2 == 1 {
			continue
		}
		for x, char := range line {
			if x%2 == 1 {
				continue
			}
			p := point.NewPoint2D(x, -y)
			if loop[p] {
				fmt.Print(string(convertChar(rune(char))))
			} else if inside[p] {
				g.Print(string(convertChar(char)))
			} else {
				r.Print(string(convertChar(char)))
			}
		}
		fmt.Println()
	}

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
