package day14

import (
	"fmt"
	"io"
	"time"
)

type info struct {
	lines []string
	pos   int
}

func rotate(lines []string) []string {
	newLines := make([]string, 0, len(lines[0]))

	for i := len(lines[0]) - 1; i >= 0; i-- {
		newLine := ""
		for _, line := range lines {
			newLine += string(line[i])
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func rotateClock(lines []string) []string {
	newLines := make([]string, 0, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		newLine := ""
		for j := len(lines) - 1; j >= 0; j-- {
			newLine += string(lines[j][i])
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func roll(lines []string) []string {
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		newLine := ""
		rounds := 0
		spaces := 0
		for _, char := range line {
			switch char {
			case '#':
				for i := 0; i < rounds; i++ {
					newLine += "O"
				}
				for i := 0; i < spaces; i++ {
					newLine += "."
				}
				newLine += "#"
				rounds = 0
				spaces = 0
			case 'O':
				rounds++
			default:
				spaces++
			}
		}
		for i := 0; i < rounds; i++ {
			newLine += "O"
		}
		for i := 0; i < spaces; i++ {
			newLine += "."
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func toString(lines []string) string {
	newLine := ""
	for _, line := range lines {
		newLine += line
	}
	return newLine
}

func cycle(i int, data info, cache map[string]info) (info, int) {
	if val, ok := cache[toString(data.lines)]; ok {
		return val, i - val.pos
	}
	newLines := make([]string, len(data.lines))
	copy(newLines, data.lines)
	for i := 0; i < 4; i++ {
		newLines = roll(newLines)
		newLines = rotateClock(newLines)
	}
	newData := info{
		lines: newLines,
		pos:   i,
	}
	cache[toString(data.lines)] = newData
	return newData, 0
}

func Part2Val(lines []string) (int, error) {
	var value int

	lines = rotate(lines)

	data := info{
		lines: lines,
		pos:   0,
	}

	cache := make(map[string]info)

	var i int
	var sinceLast int
	foundCycle := false
	cycleLen := 0
	for i = 0; i < 1000000000; i++ {
		data, sinceLast = cycle(i, data, cache)
		if !foundCycle && sinceLast > 0 {
			cycleLen = sinceLast
			foundCycle = true
			for j := i; j < 1000000000; j += cycleLen {
				i = j
			}

		}
	}

	lines = rotateClock(data.lines)

	for j, line := range lines {
		for _, char := range line {
			if char == 'O' {
				value += len(lines) - j
			}
		}
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
