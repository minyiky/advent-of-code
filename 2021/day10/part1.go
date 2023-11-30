package day10

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	corruptChars := map[rune]int{
		'[': 0,
		'(': 0,
		'{': 0,
		'<': 0,
	}

loop:
	for _, line := range lines {
		message := NewStack[rune]()
		for _, char := range line {
			switch char {
			case '[', '(', '{', '<':
				message.Push(char)
			case ']':
				if message.Pop() != '[' {
					corruptChars['[']++
					continue loop
				}
			case ')':
				if message.Pop() != '(' {
					corruptChars['(']++
					continue loop
				}
			case '}':
				if message.Pop() != '{' {
					corruptChars['{']++
					continue loop
				}
			case '>':
				if message.Pop() != '<' {
					corruptChars['<']++
					continue loop
				}
			}
		}
	}

	value = 3 * corruptChars['(']
	value += 57 * corruptChars['[']
	value += 1197 * corruptChars['{']
	value += 25137 * corruptChars['<']

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
