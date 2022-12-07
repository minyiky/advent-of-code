package day07

import (
	"log"
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

func Part2(lines []string) error {
	value, dir := Part2Val(lines)
	log.Printf("To free up enough space you should delete %s which has a total size of %d", dir, value)
	return nil
}
