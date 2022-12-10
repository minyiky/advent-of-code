package day10

import (
	"fmt"
	"io"
	"strings"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func getPixel(spritePos int, cycle int) string {
	linePos := (cycle - 1) % 40
	if aocutils.Abs(linePos-spritePos) > 1 {
		return " "
	}
	return "█"
}

func Part2Val(lines []string) string {
	var cycle int
	var rows [6]string
	x := 1
	for _, line := range lines {
		cycle++
		rows[(cycle-1)/40] += getPixel(x, cycle)

		if strings.HasPrefix(line, "noop") {
			continue
		}

		cycle++
		rows[(cycle-1)/40] += getPixel(x, cycle)

		var number int
		fmt.Sscanf(line, "addx %d", &number)

		x += number
	}

	return strings.Join(rows[:], "\n")
}

func Part2(w io.Writer, lines []string) {
	value := Part2Val(lines)
	fmt.Fprintf(w, "Reading the display, the following message was shown:\n%s\n", value)
}
