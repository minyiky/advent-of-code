package day02

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var valueVertical, valueHorizontal int

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		fmt.Sscanf(line, "%s %d", &instructions[i].direction, &instructions[i].value)
		switch instructions[i].direction {
		case "forward":
			valueHorizontal += instructions[i].value
		case "down":
			valueVertical += instructions[i].value
		case "up":
			valueVertical -= instructions[i].value
		}
	}

	return valueVertical * valueHorizontal, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, _ := Part1Val(lines)
	duration := time.Since(start)
	fmt.Fprintf(w, "The submarine moved %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
