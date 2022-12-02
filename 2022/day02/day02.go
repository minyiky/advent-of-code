package day02

import (
	_ "embed"
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

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 02 --")
	if err := Part1(lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
