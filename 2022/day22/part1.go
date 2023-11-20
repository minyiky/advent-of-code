package day22

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type Instruction struct {
	direction point.Point2D
	distance  int
}

func GetFirst(line string) point.Point2D {
	var first point.Point2D
	for j, char := range []rune(line) {
		if char == '.' {
			first = point.NewPoint2D(j+1, 1)
			break
		}
	}
	return first
}

func ExtractGrid(lines []string) (map[point.Point2D]bool, map[point.Point2D]bool) {
	grid := make(map[point.Point2D]bool)
	blocks := make(map[point.Point2D]bool)

	for i, line := range lines {
		if line == "" {
			break
		}
		for j, char := range []rune(line) {
			switch char {
			case '#':
				blocks[point.NewPoint2D(j+1, i+1)] = true
				fallthrough
			case '.':
				grid[point.NewPoint2D(j+1, i+1)] = true
			}
		}
	}

	return grid, blocks
}

func ExtractInsructions(line string) ([]Instruction, error) {
	dirList := []point.Point2D{
		point.NewPoint2D(0, 1),  // U
		point.NewPoint2D(1, 0),  // R
		point.NewPoint2D(0, -1), // D
		point.NewPoint2D(-1, 0), // L
	}

	line += "X"

	dir := 1
	var oldDir int

	instructions := make([]Instruction, 0)
	tmp := []rune{}

loop:
	for _, char := range line {
		oldDir = dir
		switch char {
		case 'R':
			dir += 3
			dir %= 4
		case 'L':
			dir += 1
			dir %= 4
		case 'X':
			break
		default:
			tmp = append(tmp, char)
			continue loop
		}

		dist, err := strconv.Atoi(string(tmp))
		if err != nil {
			return nil, err
		}
		tmp = []rune{}

		instructions = append(instructions, Instruction{
			direction: dirList[oldDir],
			distance:  dist,
		})
	}

	return instructions, nil
}

func Part1Val(lines []string) (int, error) {
	current := GetFirst(lines[0])
	grid, blocks := ExtractGrid(lines)
	instructions, err := ExtractInsructions(lines[len(lines)-1])
	if err != nil {
		return 0, err
	}

	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			next := point.Add(current, instruction.direction)
			if _, ok := grid[next]; !ok {
				rev := point.NewPoint2D(instruction.direction.X()*-1, instruction.direction.Y()*-1)
				for {
					tmpNext := point.Add(next, rev)
					if _, ok := grid[tmpNext]; !ok {
						break
					}
					next = tmpNext
				}
			}
			if blocks[next] {
				break
			}
			current = next
			continue
		}
	}

	dirList := []point.Point2D{
		point.NewPoint2D(1, 0),  // R
		point.NewPoint2D(0, 1),  // U
		point.NewPoint2D(-1, 0), // L
		point.NewPoint2D(0, -1), // D
	}

	x, _ := container.SliceContains(dirList, instructions[len(instructions)-1].direction)

	fmt.Println("Row: ", current.Y())
	fmt.Println("Col: ", current.X())
	fmt.Println("Dir: ", x)

	return 1000*current.Y() + 4*current.X() + x, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The final password was %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
