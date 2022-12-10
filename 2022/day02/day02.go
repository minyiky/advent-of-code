package day02

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

var playPoints = map[string]int{
	"X": 1, "Y": 2, "Z": 3,
	"A": 1, "B": 2, "C": 3,
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 02 --\n")
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
