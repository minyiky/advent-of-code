package day01

import (
	"log"
	"sort"
	"strconv"
)

func Part2(lines []string) error {
	elfMaxes := make(sort.IntSlice, 0, len(lines))
	var elfTotal int

	for _, line := range lines {
		if line == "" {
			elfMaxes = append(elfMaxes, elfTotal)
			elfTotal = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		elfTotal += val
	}

	sort.Sort(elfMaxes)

	sum := 0
	for _, val := range elfMaxes[len(elfMaxes)-3:] {
		sum += val
	}
	log.Printf("The top three elves carried a total of %d calories\n", sum)
	return nil
}
