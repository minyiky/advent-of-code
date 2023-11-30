package day10

import (
	"fmt"
	"io"
	"slices"
	"time"
)

func Part2Val(lines []string) (int, error) {
	values := make([]int, 0, len(lines))

loop:
	for _, line := range lines {
		message := NewStack[rune]()
		for _, char := range line {
			switch char {
			case '[', '(', '{', '<':
				message.Push(char)
			case ']':
				if message.Pop() != '[' {
					continue loop
				}
			case ')':
				if message.Pop() != '(' {
					continue loop
				}
			case '}':
				if message.Pop() != '{' {
					continue loop
				}
			case '>':
				if message.Pop() != '<' {
					continue loop
				}
			}
		}

		tempValue := 0

		for {
			char := message.Pop()
			tempValue *= 5
			switch char {
			case '[':
				tempValue += 2
			case '(':
				tempValue += 1
			case '{':
				tempValue += 3
			case '<':
				tempValue += 4
			}

			if message.IsEmpty() {
				break
			}
		}
		values = append(values, tempValue)
	}

	slices.Sort(values)

	return values[len(values)/2], nil
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
