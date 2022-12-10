package day03

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var input string

func getItems(compartment string) map[rune]bool {
	items := make(map[rune]bool)
	for _, c := range compartment {
		items[c] = true
	}
	return items
}

func getPriority(item rune) (int, error) {
	value := int(item)
	aValue, AValue := 97, 65

	if value >= aValue && value < aValue+26 {
		return value - aValue + 1, nil
	}

	if value >= AValue && value < AValue+26 {
		return value - AValue + 27, nil
	}

	return 0, errors.New("rune found outside of expected range")
}

func getMatch(compartment string, items map[rune]bool) (rune, error) {
	for _, c := range compartment {
		if _, ok := items[c]; ok {
			return c, nil
		}
	}

	var emptyRune rune
	return emptyRune, errors.New("unable to find matching character")
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 03 --\n")
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
