package day11

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Monkey struct {
	items       []uint64
	operator    string
	agent       string
	divisor     int
	trueMonkey  int
	falseMonkey int
	inspected   int
}

func NewMonkey(lines []string) *Monkey {
	monkey := &Monkey{}

	// Get starting items
	// TODO: Add error checking
	itemParts := strings.Split(lines[0], ":")
	itemStrings := strings.Split(itemParts[1], ",")
	for _, item := range itemStrings {
		num, _ := strconv.Atoi(strings.TrimSpace(item))
		monkey.items = append(monkey.items, uint64(num))
	}

	// Get Operation
	for _, operator := range "+-*/" {
		operator := string(operator)
		if strings.Contains(lines[1], operator) {
			lineSplit := strings.Split(lines[1], operator)
			monkey.agent = strings.TrimSpace(lineSplit[1])
			monkey.operator = operator
		}
	}

	// Get test
	fmt.Sscanf(lines[2], "  Test: divisible by %d", &monkey.divisor)

	// Get true monkey
	fmt.Sscanf(lines[3], "    If true: throw to monkey %d", &monkey.trueMonkey)
	// Get false monkey
	fmt.Sscanf(lines[4], "    If false: throw to monkey %d", &monkey.falseMonkey)

	return monkey
}

func (m *Monkey) Operation(x uint64) uint64 {
	var y uint64
	if m.agent == "old" {
		y = x
	} else {
		inter, _ := strconv.Atoi(m.agent)
		y = uint64(inter)
	}

	switch m.operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return 0
	}
}

func (m *Monkey) Test(x uint64) bool {
	return (x % uint64(m.divisor)) == 0
}

func worryCalculator(lines []string, divisor, rounds int) int {
	var monkeyLines []string
	var monkeys []*Monkey
	var bigDivisor uint64 = 1
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			continue
		}
		if line == "" {
			monkeys = append(monkeys, NewMonkey(monkeyLines))
			bigDivisor *= uint64(monkeys[len(monkeys)-1].divisor)
			monkeyLines = []string{}
			continue
		}
		monkeyLines = append(monkeyLines, line)
	}

	for i := 0; i < rounds; i++ {
		// if i%1000 == 0 {
		// 	fmt.Printf("== After round %d ==\n", i)
		// 	for m, monkey := range monkeys {
		// 		fmt.Printf("Monkey %d inspected items %d times.\n", m, monkey.inspected)
		// 	}
		// 	fmt.Println("")
		// }
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.inspected++
				item = monkey.Operation(item)
				item /= uint64(divisor)
				passes := monkey.Test(item)
				item %= bigDivisor
				if passes {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, item)
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, item)
				}
			}
			monkey.items = []uint64{}
		}
	}

	numInspected := make(sort.IntSlice, 0, len(monkeys))
	for _, monkey := range monkeys {
		numInspected = append(numInspected, monkey.inspected)
	}
	numInspected.Sort()

	value := numInspected[len(numInspected)-1] * numInspected[len(numInspected)-2]

	return value
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 11 --\n")
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
