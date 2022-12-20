package day19

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		Robots, maxNeeded, power := ParseLine(line)

		geodes := getMaxGeode(
			Reasources{},
			Reasources{ore: 1},
			Reasources{ore: 1},
			maxNeeded,
			Robots,
			23, 0,
		)
		value += power * geodes
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Using the blueprints, you find that you could have a maximum power of %d for extracting goedes\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
