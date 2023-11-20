package day06

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

//go:embed input.txt
var input string

func findSequence(line string, n int) (int, error) {
	marker := make([]rune, 0, n-1)

	for i, char := range line {
		index, contains := container.SliceContains(marker, char)

		if contains {
			marker = append(marker[index+1:], char)
			continue
		}

		if len(marker) < n-1 {
			marker = append(marker, char)
			continue
		}

		return i + 1, nil
	}

	return 0, errors.New("no markers found")
}

func Run(w io.Writer) {
	fmt.Fprintf(w, "-- Solution for 2022 day 06 --\n")
	if err := Part1(w, input); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, input); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
