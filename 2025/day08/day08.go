package day08

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

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2025 day 08 --\n")
	if err := Part1(w , lines); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := Part2(w , lines); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
