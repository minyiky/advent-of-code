package day03

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

func getValue(list []string, match bool) int {
	for i := 0; i < len(list[0]); i++ {
		if len(list) == 1 {
			break
		}

		commonBit := 0

		for _, line := range list {
			if line[i:i+1] == "1" {
				commonBit++
			} else {
				commonBit--
			}
		}

		if commonBit < 0 {
			commonBit = 0
		} else {
			commonBit = 1
		}

		newList := make([]string, 0, len(list))
		for _, oxy := range list {
			if (oxy[i:i+1] == strconv.Itoa(commonBit)) == match {
				newList = append(newList, oxy)
			}
		}
		list = newList
	}

	var oxygen int

	for i, bit := range list[0] {
		if bit > 48 {
			oxygen |= 1 << (len(list[0]) - i - 1)
		}
	}

	return oxygen
}

func Part2Val(lines []string) (int, error) {
	oxygen := getValue(container.CopySlice(lines), true)
	carbon := getValue(container.CopySlice(lines), false)

	value := oxygen * carbon

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
