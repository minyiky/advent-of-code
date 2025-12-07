package day01

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	return Part1ValWithViz(lines, nil, 1.0)
}

// Part1ValWithViz runs Part1 with visualization support
func Part1ValWithViz(lines []string, vizCallback func(data map[string]interface{}), speed float64) (int, error) {
	var value int
	lockValue := 50

	// Send initial state
	if vizCallback != nil {
		vizCallback(map[string]interface{}{
			"position":         lockValue,
			"counter":          value,
			"instructionIndex": -1,
		})
	}

	for instrIdx, line := range lines {
		var dir rune
		var num int
		if _, err := fmt.Sscanf(line, "%c%d", &dir, &num); err != nil {
			return 0, err
		}

		// Calculate final position for this move
		finalLockValue := lockValue
		switch dir {
		case 'L':
			finalLockValue -= num
		case 'R':
			finalLockValue += num
		}
		finalLockValue %= 100
		if finalLockValue < 0 {
			finalLockValue += 100
		}

		// Animate the rotation click by click
		for i := 0; i < num; i++ {
			switch dir {
			case 'L':
				lockValue -= 1
			case 'R':
				lockValue += 1
			}

			lockValue %= 100
			if lockValue < 0 {
				lockValue += 100
			}

			if lockValue == 0 {
				value += 1
			}

			// Send visualization update for each click
			if vizCallback != nil {
				vizCallback(map[string]interface{}{
					"position":         lockValue,
					"finalPosition":    finalLockValue,
					"counter":          value,
					"instructionIndex": instrIdx,
				})
			}
		}
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
