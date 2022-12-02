package day01

import (
	"log"
	"sort"
	"strconv"
)

func Part2Val(lines []string) (int, error) {
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
			return 0, err
		}
		elfTotal += val
	}
	if elfTotal != 0 {
		elfMaxes = append(elfMaxes, elfTotal)
	}

	sort.Sort(elfMaxes)

	sum := 0
	for _, val := range elfMaxes[len(elfMaxes)-3:] {
		sum += val
	}

	return sum, nil
}

func Part2(lines []string) error {
	sum, err := Part2Val(lines)
	if err != nil {
		return err
	}
	log.Printf("The top three elves carried a total of %d calories\n", sum)
	return nil
}
