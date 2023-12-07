package day07

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

var replacer = strings.NewReplacer(
	"2", "a",
	"3", "b",
	"4", "c",
	"5", "d",
	"6", "e",
	"7", "f",
	"8", "g",
	"9", "h",
	"T", "i",
	"J", "j",
	"Q", "k",
	"K", "l",
	"A", "m",
)

var replacerJack = strings.NewReplacer(
	"2", "a",
	"3", "b",
	"4", "c",
	"5", "d",
	"6", "e",
	"7", "f",
	"8", "g",
	"9", "h",
	"T", "i",
	"Q", "k",
	"K", "l",
	"A", "m",
)

type Hand struct {
	hand  string
	score int
	bid   int
}

func NewHand(line string) Hand {
	fields := strings.Fields(line)
	hand := replacer.Replace(fields[0])
	chars := make(map[rune]int)
	for _, c := range hand {
		if _, ok := chars[c]; !ok {
			chars[c] = 0
		}
		chars[c]++
	}

	nums := make([]int, 0)
	for _, v := range chars {
		nums = append(nums, v)
	}

	slices.Sort(nums)

	var score int

	switch {
	case nums[0] == 5:
		score = 6
	case nums[1] == 4:
		score = 5
	case nums[1] == 3:
		score = 4
	case nums[2] == 3:
		score = 3
	case nums[2] == 2:
		score = 2
	case nums[3] == 2:
		score = 1
	}

	bid, _ := strconv.Atoi(fields[1])

	return Hand{
		hand:  hand,
		score: score,
		bid:   bid,
	}
}

func NewJackHand(line string) Hand {
	fields := strings.Fields(line)
	hand := replacerJack.Replace(fields[0])
	chars := make(map[rune]int)
	jacks := 0
	for _, c := range hand {
		if c == 'J' {
			jacks++
			continue
		}
		if _, ok := chars[c]; !ok {
			chars[c] = 0
		}
		chars[c]++
	}

	nums := make([]int, 0)
	for _, v := range chars {
		nums = append(nums, v)
	}

	slices.Sort(nums)

	if jacks != 5 {
		nums[len(nums)-1] += jacks
	}

	var score int

	switch {
	case len(nums) == 0: // JJJJJ
		score = 6
	case nums[0] == 5: // AAAAA [5]
		score = 6
	case nums[1] == 4: // AAAA. [1, 4]
		score = 5
	case nums[1] == 3: // AAABB [2, 3]
		score = 4
	case nums[2] == 3: // AAA.. [1, 1, 3]
		score = 3
	case nums[2] == 2: // AABB. [1, 2, 2]
		score = 2
	case nums[3] == 2: // AA... [1, 1, 1, 2]
		score = 1
	}

	bid, _ := strconv.Atoi(fields[1])

	return Hand{
		hand:  hand,
		score: score,
		bid:   bid,
	}
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2023 day 07 --\n")
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
