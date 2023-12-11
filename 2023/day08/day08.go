package day08

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("([1-9A-Z]+)")

type Node struct {
	paths map[rune]string
}

//go:embed input.txt
var input string

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2023 day 08 --\n")
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
