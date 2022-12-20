package day19

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
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
	return r.ore >= t.ore && r.clay >= t.clay && r.obsidian >= t.obsidian
}

type Robot struct {
	cost     Reasources
	produces Reasources
}

func mid(stock, gen, newGen, maxNeeded Reasources, robots []Robot, time, geodes, rNum int) int {
	robot := robots[rNum]
	var extraTime int
	tmpStock := stock
	for !tmpStock.IsEnough(robot.cost) {
		tmpStock = tmpStock.Add(gen)
		extraTime++
	}

	if extraTime+1 >= time {
		return geodes
	}
	// Take the reasources needed from the stockpile
	tmpStock = tmpStock.Sub(robot.cost)

	// Produce the robot by adding to the relevent generator
	tmpGen := gen.Add(robot.produces)

	return getMaxGeode(tmpStock, gen, tmpGen, maxNeeded, robots, time-1-extraTime, geodes)
}

func getMaxGeode(stock, gen, newGen, maxNeeded Reasources, robots []Robot, time, geodes int) int {
	stock = stock.Add(gen)

	gen = newGen

	if time == 0 {
		return stock.geode
	}

	if time == 1 {
		return getMaxGeode(stock, gen, gen, maxNeeded, robots, time-1, geodes)
	}

	tmpGeodes := geodes

	// Make a Geode
	if gen.obsidian > 0 {
		geode := mid(stock, gen, newGen, maxNeeded, robots, time, tmpGeodes, 0)
		if geode > geodes {
			geodes = geode
		}
	}

	// Make an Obsidian
	if gen.clay > 0 && gen.obsidian < maxNeeded.obsidian {
		geode := mid(stock, gen, newGen, maxNeeded, robots, time, tmpGeodes, 1)
		if geode > geodes {
			geodes = geode
		}
	}

	// Make a clay
	if gen.clay < maxNeeded.clay {
		geode := mid(stock, gen, newGen, maxNeeded, robots, time, tmpGeodes, 2)
		if geode > geodes {
			geodes = geode
		}
	}

	// Make an ore
	if gen.ore < maxNeeded.ore {
		geode := mid(stock, gen, newGen, maxNeeded, robots, time, tmpGeodes, 3)
		if geode > geodes {
			geodes = geode
		}
	}

	geode := stock.geode + gen.geode*time
	if geode > geodes {
		geodes = geode
	}

	return geodes
}

func ParseLine(line string) ([]Robot, Reasources, int) {
	var power, oreBotOre, clayBotOre, obsBotOre, obsBotClay, geoBotOre, geoBotObs int
	fmt.Sscanf(
		line,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&power, &oreBotOre, &clayBotOre, &obsBotOre, &obsBotClay, &geoBotOre, &geoBotObs,
	)

	oreList := sort.IntSlice{oreBotOre, clayBotOre, obsBotOre, geoBotOre}
	sort.Sort(oreList)
	maxNeeded := Reasources{
		ore:      oreList[3],
		clay:     obsBotClay,
		obsidian: geoBotObs,
		geode:    math.MaxInt,
	}

	Robots := make([]Robot, 4)
	Robots[3] = Robot{
		cost:     Reasources{ore: oreBotOre},
		produces: Reasources{ore: 1},
	}
	Robots[2] = Robot{
		cost:     Reasources{ore: clayBotOre},
		produces: Reasources{clay: 1},
	}
	Robots[1] = Robot{
		cost:     Reasources{ore: obsBotOre, clay: obsBotClay},
		produces: Reasources{obsidian: 1},
	}
	Robots[0] = Robot{
		cost:     Reasources{ore: geoBotOre, obsidian: geoBotObs},
		produces: Reasources{geode: 1},
	}
	return Robots, maxNeeded, power
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
