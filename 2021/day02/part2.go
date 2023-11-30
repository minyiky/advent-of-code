package day02

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var valueVertical, valueHorizontal, aim int

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		fmt.Sscanf(line, "%s %d", &instructions[i].direction, &instructions[i].value)
		switch instructions[i].direction {
		case "forward":
			valueHorizontal += instructions[i].value
			valueVertical += aim * instructions[i].value
		case "down":
			aim += instructions[i].value
		case "up":
			aim -= instructions[i].value
		}
	}

	return valueVertical * valueHorizontal, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "With the new info the submarine moved %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
