package day01

import (
	"fmt"
	"io"
	"time"
)

func Part2Val(lines []string) (int, error) {
	return Part2ValWithViz(lines, nil, 1.0)
}

// Part2ValWithViz runs Part2 with visualization support using factory pattern
func Part2ValWithViz(lines []string, vizFactory func(setup map[string]interface{}) func(data map[string]interface{}), speed float64) (int, error) {
	var value int
	lockValue := 50

	// Initialize visualization with setup data
	var vizFn func(data map[string]interface{})
	if vizFactory != nil {
		vizFn = vizFactory(map[string]interface{}{
			"type":         "dial",
			"maxValue":     100,
			"instructions": lines,
		})
	}

	// Send initial state
	if vizFn != nil {
		vizFn(map[string]interface{}{
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
			prevLockValue := lockValue

			switch dir {
			case 'L':
				lockValue -= 1
			case 'R':
				lockValue += 1
			}

			// Check for complete rotations on each step
			if lockValue >= 100 || lockValue <= -100 {
				value += func() int {
					v := lockValue / 100
					if v > 0 {
						return v
					}
					return -1 * v
				}()
			}

			// Check for crossing zero
			if lockValue < 0 && prevLockValue != 0 {
				value += 1
			}

			if lockValue == 0 {
				value += 1
			}

			// Normalize to 0-99 range
			lockValue %= 100
			if lockValue < 0 {
				lockValue += 100
			}

			// Send visualization update for each click
			if vizFn != nil {
				vizFn(map[string]interface{}{
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
