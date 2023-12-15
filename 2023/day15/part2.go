package day15

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type lens struct {
	label string
	value int
}

func Part2Val(lines []string) (int, error) {
	var value int

	line := lines[0]

	boxes := make([][]lens, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = make([]lens, 0)
	}

parts:
	for _, part := range strings.Split(line, ",") {
		box := 0
		for _, char := range part {
			if char == '=' || char == '-' {
				break
			}
			box += int(char)
			box *= 17
			box %= 256
		}
		if strings.Contains(part, "=") {
			label, value, _ := strings.Cut(part, "=")
			v, _ := strconv.Atoi(value)
			for i := range boxes[box] {
				if boxes[box][i].label == label {
					boxes[box][i] = lens{label: label, value: v}
					continue parts
				}
			}
			boxes[box] = append(boxes[box], lens{label: label, value: v})
		} else {
			newLenses := make([]lens, 0, len(boxes[box]))
			label, _, _ := strings.Cut(part, "-")
			for _, lens := range boxes[box] {
				if lens.label == label {
					continue
				}
				newLenses = append(newLenses, lens)
			}
			boxes[box] = newLenses
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			value += lens.value * (i + 1) * (j + 1)
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
