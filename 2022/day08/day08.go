package day08

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 08 --")
	Part1(lines)
	Part2(lines)
}
