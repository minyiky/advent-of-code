package day19

import (
	"fmt"
	"io"
	"math"
	"sort"
	"time"
)

func realistic(gen, stock Reasources, timeLeft, target int) bool {
	for i := 1; i <= timeLeft; i++ {
		gen.geode += 1
		stock.geode += gen.geode
	}
	return stock.geode >= target
}

func getMaxGeode(stock, gen Reasources, maxNeeded []int, robots []Robot, time, geodes int) int {
	stock = stock.Add(gen)

	if time == 24 {
		return stock.geode
	}

	if time == 23 {
		return getMaxGeode(stock, gen, maxNeeded, robots, time+1, geodes)
	}

	tmpGeodes := geodes

	resList := []int{
		0, gen.obsidian, gen.clay, gen.ore,
	}
	for i, robot := range robots {
		if !stock.IsEnough(robot.cost) {
			continue
		}

		if resList[i] >= maxNeeded[i] {
			continue
		}

		tmpStock := stock.Sub(robot.cost)
		tmpStock = tmpStock.Add(gen)
		tmpGen := gen.Add(robot.produces)
		geode := getMaxGeode(tmpStock, tmpGen, maxNeeded, robots, time+2, tmpGeodes)
		if geode > geodes {
			geodes = geode
		}
		if geode == 6 {
			return 6
		}
	}

	geode := getMaxGeode(stock, gen, maxNeeded, robots, time+1, tmpGeodes)
	if geode > geodes {
		geodes = geode
	}

	return geodes
}

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		var num, oreBotOre, clayBotOre, obsBotOre, obsBotClay, geoBotOre, geoBotObs int
		fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&num, &oreBotOre, &clayBotOre, &obsBotOre, &obsBotClay, &geoBotOre, &geoBotObs,
		)

		oreList := sort.IntSlice{oreBotOre, clayBotOre, obsBotOre, geoBotOre}
		sort.Sort(oreList)
		maxNeeded := []int{
			math.MaxInt,
			geoBotObs,
			obsBotClay,
			oreList[3],
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

		geodes := getMaxGeode(
			Reasources{},
			Reasources{ore: 1},
			maxNeeded,
			Robots,
			1, 0,
		)
		value += num * geodes
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
