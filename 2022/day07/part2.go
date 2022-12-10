package day07

import (
	"fmt"
	"io"
	"math"
)

func Part2Val(lines []string) (int, string) {
	dirs := ReadDirs(lines)

	spaceNeeded := 30000000 - (70000000 - dirs["/"])
	var dir string
	minToDel := math.MaxInt

	for k, v := range dirs {
		if (spaceNeeded <= v) && (v < minToDel) {
			dir = k
			minToDel = v
		}
	}

	return minToDel, dir
}

func Part2(w io.Writer, lines []string) {
	value, dir := Part2Val(lines)
	fmt.Fprintf(w, "To free up enough space you should delete %s which has a total size of %d\n", dir, value)
}
