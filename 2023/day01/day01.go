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

func calculate(line string) int {
	var value int

	code := make([]rune, 0)
	for _, char := range line {
		if char-'0' >= 0 && char-'0' <= 9 {
			code = append(code, char)
		}
	}
	value += int(code[0]-'0')*10 + int(code[len(code)-1]-'0')

	return value
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
	fmt.Println("")
}
