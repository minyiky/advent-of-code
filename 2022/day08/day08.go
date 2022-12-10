package day08

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed input.txt
var input string

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 08 --\n")
	Part1(w, lines)
	Part2(w, lines)
}
