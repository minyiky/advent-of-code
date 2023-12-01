package day01

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func calculate(lines []string) (int, error) {
	var value int

	codes := make([][]rune, len(lines))

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, line := range lines {
		for i, word := range words {
			line = strings.Replace(line, word, fmt.Sprintf("%d", i), -1)
		}
		for _, char := range line {
			if char-'0' >= 0 && char-'0' <= 9 {
				codes[i] = append(codes[i], char)
			}
		}
	}

	for _, code := range codes {
		value += int(code[0]-'0')*10 + int(code[len(code)-1]-'0')
	}

	return value, nil
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2023 day 01 --\n")
	if err := Part1(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
