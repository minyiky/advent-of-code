package day07

import (
	"fmt"
	"io"
)

func Part1Val(lines []string) int {
	var value int

	dirs := ReadDirs(lines)

	for _, v := range dirs {
		if v < 100000 {
			value += v
		}
	}

	return value
}

func Part1(w io.Writer, lines []string) {
	value := Part1Val(lines)
	fmt.Fprintf(w, "The total size of directories under 100000 was %d\n", value)
}
