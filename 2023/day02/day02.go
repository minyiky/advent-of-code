package day02

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
	rRed   = regexp.MustCompile("([0-9]+) red")
	rGreen = regexp.MustCompile("([0-9]+) green")
	rBlue  = regexp.MustCompile("([0-9]+) blue")
)

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2023 day 02 --\n")
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
