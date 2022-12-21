package day21

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Monkey struct {
	Number    int
	Known     bool
	Opperator rune
	ReliesOn  []string
}

func NewMonkey(line string) (string, *Monkey) {
	monkey := &Monkey{}
	name := line[:4]
	line = line[6:]
	num, err := strconv.Atoi(line)
	if err != nil {
		monkey.ReliesOn = []string{line[:4], line[7:]}
		monkey.Opperator = []rune(line)[5]
	} else {
		monkey.Known = true
		monkey.Number = num
	}

	return name, monkey
}

func (m *Monkey) GetVal(monkeyMap map[string]*Monkey) int {
	if m.Known {
		return m.Number
	}

	v0 := monkeyMap[m.ReliesOn[0]].GetVal(monkeyMap)
	v1 := monkeyMap[m.ReliesOn[1]].GetVal(monkeyMap)

	switch m.Opperator {
	case '+':
		m.Number = v0 + v1
	case '-':
		m.Number = v0 - v1
	case '*':
		m.Number = v0 * v1
	case '/':
		m.Number = v0 / v1
	}
	return m.Number
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 21 --\n")
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
