package day04

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getRange(task string) ([]int, error) {
	ends := strings.Split(task, "-")
	if len(ends) != 2 {
		return nil, errors.New("wrong number of sections found")
	}

	endVals := make([]int, 2)
	for i, end := range ends {
		val, err := strconv.Atoi(end)
		if err != nil {
			return nil, err
		}

		endVals[i] = val
	}

	return endVals, nil
}

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 04 --")
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
