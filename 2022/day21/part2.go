package day21

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	monkeys := make(map[string]*Monkey)

	for _, line := range lines {
		name, monkey := NewMonkey(line)
		monkeys[name] = monkey
	}

	// Check the branch to track
	monkeys["humn"].Known = true
	monkeys["humn"].Number = 0
	v1 := monkeys[monkeys["root"].ReliesOn[0]].GetVal(monkeys)
	monkeys["humn"].Number = 10000
	v2 := monkeys[monkeys["root"].ReliesOn[0]].GetVal(monkeys)

	var boi int
	if v1 == v2 {
		boi = 1
		monkeys["humn"].Number = 0
		v1 = monkeys[monkeys["root"].ReliesOn[1]].GetVal(monkeys)
		monkeys["humn"].Number = 10000
		v2 = monkeys[monkeys["root"].ReliesOn[1]].GetVal(monkeys)
	}

	mult := 1
	if v2 > v1 {
		mult = -1
	}

	refVal := monkeys[monkeys["root"].ReliesOn[1-boi]].GetVal(monkeys)
	moi := monkeys[monkeys["root"].ReliesOn[boi]]

	value, ok := sort.Find(1e16, func(v int) int {
		monkeys["humn"].Number = v
		return mult*moi.GetVal(monkeys) - mult*refVal
	})
	if !ok {
		return 0, errors.New("unable to find valid value")
	}

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Those elephants really are silly arent they, now that their translation is correct you realise you should hout out %d for root's values to be the same\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
