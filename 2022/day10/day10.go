package day10

import (
	_ "embed"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 10 --")
	if err := Part1(lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	Part2(lines)
}
