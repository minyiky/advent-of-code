package day01

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")
	ints := make([]int, len(lines))
	for i, line := range lines {
		var err error
		ints[i], err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(w, "-- Solution for 2021 day 01 --\n")
	if err := Part1(w, ints); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, ints); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
