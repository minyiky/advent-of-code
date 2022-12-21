package day21

import (
	"fmt"
	"io"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	monkeys := make(map[string]*Monkey)

	for _, line := range lines {
		name, monkey := NewMonkey(line)
		monkeys[name] = monkey
	}

	value = monkeys["root"].GetVal(monkeys)

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "You realise that the monkey named 'root' is going to shout out %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
