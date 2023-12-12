package day12

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

var (
	rNum   = regexp.MustCompile(`([\d]+)`)
	rCog   = regexp.MustCompile(`(#+)`)
	rQMark = regexp.MustCompile(`(\?)`)
)

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2023 day 12 --\n")
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
