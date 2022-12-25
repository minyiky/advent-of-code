package day22

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

var dirList = []aocutils.Vector{
	aocutils.NewVector(1, 0),  // R
	aocutils.NewVector(0, 1),  // U
	aocutils.NewVector(-1, 0), // L
	aocutils.NewVector(0, -1), // D
}

func makeCrossBoundry(size int) func(aocutils.Vector, int) (aocutils.Vector, int, error) {
	return func(pos aocutils.Vector, dir int) (aocutils.Vector, int, error) {
		if pos.X <= 0 {
			pos.X -= size
		}
		if pos.Y <= 0 {
			pos.Y -= size
		}
		markX := (pos.X - 1) / size
		markY := (pos.Y - 1) / size

		if markX == 0 && markY == 0 && dir == 3 {
			newPos := aocutils.NewVector(3*size-((pos.X-1)%size), 1)
			newdir := 1
			return newPos, newdir, nil
		}
		if markX == 2 && markY == -1 && dir == 3 {
			newPos := aocutils.NewVector(size-((pos.X-1)%size), size+1)
			newdir := 1
			return newPos, newdir, nil
		}

		if markX == 0 && markY == 2 && dir == 1 {
			newPos := aocutils.NewVector(3*size-((pos.X-1)%size), 3*size)
			newdir := 3
			return newPos, newdir, nil
		}
		if markX == 2 && markY == 3 && dir == 1 {
			newPos := aocutils.NewVector(size-((pos.X-1)%size), 2*size)
			newdir := 3
			return newPos, newdir, nil
		}

		if markX == -1 && markY == 1 && dir == 2 {
			newPos := aocutils.NewVector(4*size-((pos.Y-1)%size), 3*size)
			newdir := 3
			return newPos, newdir, nil
		}
		if markX == 3 && markY == 3 && dir == 1 {
			newPos := aocutils.NewVector(1, 2*size-((pos.X-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		if markX == 3 && markY == 0 && dir == 0 {
			newPos := aocutils.NewVector(4*size, 3*size-((pos.Y-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}
		if markX == 4 && markY == 2 && dir == 0 {
			newPos := aocutils.NewVector(3*size, size-((pos.Y-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}

		if markX == 3 && markY == 1 && dir == 0 {
			newPos := aocutils.NewVector(4*size-((pos.Y-1)%size), 2*size+1)
			newdir := 1
			return newPos, newdir, nil
		}
		if markX == 3 && markY == 1 && dir == 3 {
			newPos := aocutils.NewVector(3*size, 2*size-((pos.X-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}

		if markX == 1 && markY == 0 && dir == 2 {
			newPos := aocutils.NewVector(size+1+((pos.Y-1)%size), size+1)
			newdir := 1
			return newPos, newdir, nil
		}
		if markX == 1 && markY == 0 && dir == 3 {
			newPos := aocutils.NewVector(2*size+1, 1+((pos.X-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		if markX == 1 && markY == 2 && dir == 2 {
			newPos := aocutils.NewVector(2*size-((pos.Y-1)%size), 2*size)
			newdir := 3
			return newPos, newdir, nil
		}
		if markX == 1 && markY == 2 && dir == 1 {
			newPos := aocutils.NewVector(2*size+1, 3*size-((pos.X-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		return aocutils.NewVector(0, 0), 0, errors.New("unknown grid position")
	}
}

func makeCrossBoundryReal(size int) func(aocutils.Vector, int) (aocutils.Vector, int, error) {
	return func(pos aocutils.Vector, dir int) (aocutils.Vector, int, error) {
		if pos.X <= 0 {
			pos.X -= size
		}
		if pos.Y <= 0 {
			pos.Y -= size
		}
		markX := (pos.X - 1) / size
		markY := (pos.Y - 1) / size

		if markX == 0 && markY == 0 && dir == 2 {
			newPos := aocutils.NewVector(1, 3*size-((pos.Y-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}
		if markX == -1 && markY == 2 && dir == 2 {
			newPos := aocutils.NewVector(size+1, size-((pos.Y-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		if markX == 0 && markY == 1 && dir == 2 {
			newPos := aocutils.NewVector(1+((pos.Y-1)%size), 2*size+1)
			newdir := 1
			return newPos, newdir, nil
		}
		if markX == 0 && markY == 1 && dir == 3 {
			newPos := aocutils.NewVector(size+1, size+1+((pos.X-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		if markX == -1 && markY == 3 && dir == 2 {
			newPos := aocutils.NewVector(size+1+((pos.Y-1)%size), 1)
			newdir := 1
			return newPos, newdir, nil
		}
		if markX == 1 && markY == -1 && dir == 3 {
			newPos := aocutils.NewVector(1, 3*size+1+((pos.X-1)%size))
			newdir := 0
			return newPos, newdir, nil
		}

		// Look at
		if markX == 2 && markY == -1 && dir == 3 {
			newPos := aocutils.NewVector(2*size+1+((pos.X-1)%size), 1)
			newdir := 3
			return newPos, newdir, nil
		}
		if markX == 0 && markY == 4 && dir == 1 {
			newPos := aocutils.NewVector(1+((pos.X-1)%size), 4*size)
			newdir := 1
			return newPos, newdir, nil
		}

		if markX == 1 && markY == 3 && dir == 1 {
			newPos := aocutils.NewVector(size, 3*size+1+((pos.X-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}
		if markX == 1 && markY == 3 && dir == 0 {
			newPos := aocutils.NewVector(size+1+((pos.Y-1)%size), 3*size)
			newdir := 3
			return newPos, newdir, nil
		}

		if markX == 2 && markY == 1 && dir == 1 {
			newPos := aocutils.NewVector(2*size, size+1+((pos.X-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}
		if markX == 2 && markY == 1 && dir == 0 {
			newPos := aocutils.NewVector(2*size+1+((pos.Y-1)%size), size)
			newdir := 3
			return newPos, newdir, nil
		}

		if markX == 3 && markY == 0 && dir == 0 {
			newPos := aocutils.NewVector(2*size, 3*size-((pos.Y-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}
		if markX == 2 && markY == 2 && dir == 0 {
			newPos := aocutils.NewVector(3*size, size-((pos.Y-1)%size))
			newdir := 2
			return newPos, newdir, nil
		}

		return aocutils.NewVector(0, 0), 0, errors.New("unknown grid position")
	}
}

func Part2Val(lines []string) (int, error) {
	current := GetFirst(lines[0])
	grid, blocks := ExtractGrid(lines)
	var dir int

	crossBoundry := makeCrossBoundry(len(lines[0]) / 4)

	instructionLine := lines[len(lines)-1] + "X"

	tmp := []rune{}

loop:
	for _, char := range instructionLine {
		var modifier int
		switch char {
		case 'R':
			modifier = 1
		case 'L':
			modifier = 3
		case 'X':
			modifier = 0
		default:
			tmp = append(tmp, char)
			continue loop
		}

		numToMove, err := strconv.Atoi(string(tmp))
		if err != nil {
			return 0, err
		}

		for i := 0; i < numToMove; i++ {
			next := current.Add(dirList[dir])
			tmpDir := dir
			if _, ok := grid[next]; !ok {
				next, tmpDir, err = crossBoundry(next, dir)
				if err != nil {
					return 0, err
				}
			}
			if blocks[next] {
				break
			}
			dir = tmpDir
			current = next
			continue
		}
		tmp = []rune{}
		dir += modifier
		dir %= 4

	}

	fmt.Println("Row: ", current.Y)
	fmt.Println("Col: ", current.X)
	fmt.Println("Dir: ", dir)

	return 1000*current.Y + 4*current.X + dir, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Using a cube mapping, the final password was %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
