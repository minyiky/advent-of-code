package day08

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

var letters = map[int][]rune{
	0: {'a', 'b', 'c', 'e', 'f', 'g'},
	1: {'c', 'f'},
	2: {'a', 'c', 'd', 'e', 'g'},
	3: {'a', 'c', 'd', 'f', 'g'},
	4: {'b', 'c', 'd', 'f'},
	5: {'a', 'b', 'd', 'f', 'g'},
	6: {'a', 'b', 'd', 'e', 'f', 'g'},
	7: {'a', 'c', 'f'},
	8: {'a', 'b', 'c', 'd', 'e', 'f', 'g'},
	9: {'a', 'b', 'c', 'd', 'f', 'g'},
}

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		digits := make([]string, 10)
		counts := map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
			'd': 0,
			'e': 0,
			'f': 0,
			'g': 0,
		}
		segments := make(map[rune]rune)
		mapping, disp, _ := strings.Cut(line, "|")
		for _, digit := range strings.Fields(mapping) {
			switch len(digit) {
			case 4:
				digits[4] = digit
			}
			for _, char := range digit {
				counts[char]++
			}
		}
		for char, count := range counts {
			switch count {
			case 4:
				segments[char] = 'e'
			case 6:
				segments[char] = 'b'
			case 9:
				segments[char] = 'f'
			case 8:
				if !strings.Contains(digits[4], string(char)) {
					segments[char] = 'a'
				} else {
					segments[char] = 'c'
				}
			case 7:
				if !strings.Contains(digits[4], string(char)) {
					segments[char] = 'g'
				} else {
					segments[char] = 'd'
				}
			}
		}

		midValue := 0
		for _, digit := range strings.Fields(disp) {
			var tempValue int
			switch len(digit) {
			case 2:
				tempValue = 1
			case 3:
				tempValue = 7
			case 4:
				tempValue = 4
			case 7:
				tempValue = 8
			case 5:
			case5:
				for _, i := range []int{2, 3, 5} {
					for _, char := range digit {
						if _, found := container.SliceContains(letters[i], segments[char]); !found {
							continue case5
						}
					}
					tempValue = i
					break
				}
			case 6:
			case6:
				for _, i := range []int{0, 6, 9} {
					for _, char := range digit {
						if _, found := container.SliceContains(letters[i], segments[char]); !found {
							continue case6
						}
					}
					tempValue = i
					break
				}
			}
			midValue *= 10
			midValue += tempValue
		}
		value += midValue
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
