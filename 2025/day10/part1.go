package day10

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		parts := strings.Split(line, " ")

		target := make([]bool, len(parts[0])-2)
		for i, char := range parts[0][1 : len(parts[0])-1] {
			if char == '#' {
				target[i] = true
			}
		}

		instructions := make([][]int, 0, len(parts)-2)
		for _, p := range parts[1 : len(parts)-1] {
			var instruction []int
			for _, c := range p {
				if c >= '0' && c <= '9' {
					instruction = append(instruction, int(c-'0'))
				}
			}
			if len(instruction) > 0 {
				instructions = append(instructions, instruction)
			}

		}

		states := make(map[string][]bool)

		value += func() int {
			for i := range len(instructions) {
				perms := KCombinations(instructions, i+1)
			permLoop:
				for _, perm := range perms {
					if _, ok := states[key(perm)]; ok {
						continue
					}
					state := make([]bool, len(target))
					if len(perm) > 1 {
						copy(state, states[key(perm[:len(perm)-1])])
					}

					for _, bulb := range perm[len(perm)-1] {
						state[bulb] = !state[bulb]
					}

					states[key(perm)] = state

					for i := range state {
						if state[i] != target[i] {
							continue permLoop
						}
					}

					return len(perm)
				}
			}
			return 0
		}()
	}

	return value, nil
}

// KCombinations generates all k-length combinations from arr
func KCombinations[T any](arr []T, k int) [][]T {
	var result [][]T
	n := len(arr)
	if k > n || k <= 0 {
		return result
	}

	var backtrack func(start int, current []T)
	backtrack = func(start int, current []T) {
		if len(current) == k {
			comb := make([]T, k)
			copy(comb, current)
			result = append(result, comb)
			return
		}
		for i := start; i < n; i++ {
			backtrack(i+1, append(current, arr[i]))
		}
	}
	backtrack(0, []T{})
	return result
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
