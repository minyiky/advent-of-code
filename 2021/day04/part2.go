package day04

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	numbers := strings.Split(lines[0], ",")

	boards := make([]*board, 0)

	lineSet := make([]string, 0)

	for _, line := range lines[2:] {
		if line != "" {
			lineSet = append(lineSet, line)
			continue
		}
		board, err := newBoard(lineSet)
		if err != nil {
			return 0, err
		}
		boards = append(boards, board)
		lineSet = make([]string, 0)
	}

	if len(lineSet) > 0 {
		board, err := newBoard(lineSet)
		if err != nil {
			return 0, err
		}
		boards = append(boards, board)
	}

	boardsNeeded := len(boards)

	var completeBoards int
	for _, call := range numbers {
		for _, board := range boards {
			if board.complete {
				continue
			}
			board.mark(call)
			if board.complete {
				completeBoards++
				number, err := strconv.Atoi(call)
				if err != nil {
					return 0, err
				}
				if completeBoards == boardsNeeded {
					value += number * board.score()
					break
				}
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
