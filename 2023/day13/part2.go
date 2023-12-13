package day13

import (
	"fmt"
	"io"
	"time"
)

func calcBlockSmudge(block []string) int {
	calc := func(block []string) int {
		for i := range block[:len(block)-1] {
			diffs := 0
			// fmt.Println()
			for j := 0; j < len(block)-1; j++ {
				if i-j < 0 || i+j+1 == len(block) {
					// fmt.Println(i, j, "breaking")
					break
				}
				// fmt.Println(j, block[i-j])
				// fmt.Println(j, block[j+i+1])
				upper := block[i-j]
				lower := block[i+j+1]
				for k := range block[0] {
					if upper[k] != lower[k] {
						diffs++
					}
				}
			}
			// fmt.Println(i)
			if diffs == 1 {
				return i + 1
			}
		}
		return 0
	}

	if v := calc(block); v != 0 {
		return 100 * v
	}

	newBlock := make([]string, 0, len(block[0]))

	for i := range block[0] {
		l := ""
		for _, line := range block {
			l += line[i : i+1]
		}
		newBlock = append(newBlock, l)
	}

	return calc(newBlock)
}

func Part2Val(lines []string) (int, error) {
	var value int

	block := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			value += calcBlockSmudge(block)
			block = make([]string, 0)
			continue
		}
		block = append(block, line)
	}
	value += calcBlockSmudge(block)

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
