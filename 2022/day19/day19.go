package day19

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

type Reasources struct {
	ore, clay, obsidian, geode int
}

func (r Reasources) Add(t Reasources) Reasources {
	return Reasources{
		ore:      r.ore + t.ore,
		clay:     r.clay + t.clay,
		obsidian: r.obsidian + t.obsidian,
		geode:    r.geode + t.geode,
	}
}

func (r Reasources) Sub(t Reasources) Reasources {
	return Reasources{
		ore:      r.ore - t.ore,
		clay:     r.clay - t.clay,
		obsidian: r.obsidian - t.obsidian,
		geode:    r.geode - t.geode,
	}
}

func (r Reasources) IsEnough(t Reasources) bool {
	return r.ore >= t.ore && r.clay >= t.clay && r.obsidian >= t.obsidian && r.geode >= t.geode
}

type Robot struct {
	cost     Reasources
	produces Reasources
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 19 --\n")
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
