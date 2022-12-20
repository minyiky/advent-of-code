package day19

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	value := 1

	values := []int{}

	for i, line := range lines {
		Robots, maxNeeded, _ := ParseLine(line)

		geodes := getMaxGeode(
			Reasources{},
			Reasources{ore: 1},
			Reasources{ore: 1},
			maxNeeded,
			Robots,
			31, 0,
		)
		values = append(values, geodes)
		if i == 2 {
			break
		}
	}

	for _, v := range values {
		value *= v
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
	fmt.Fprintf(w, "Even if the elephants did lose most of the blueprints, you now have food so the extra time allows you to have a power of %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
