package dayone

import (
	"log"
	"os"
	"sort"
	"strconv"
)

type TopThree []int

func (t TopThree) Len() int {
	return len(t)
}

func (t TopThree) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TopThree) Less(i, j int) bool {
	return t[i] < t[j]
}

func tryInsert(top3 TopThree, val int) TopThree {
	if len(top3) < 3 {
		top3 = append(top3, val)
		sort.Sort(top3)
		return top3
	}

	if val < top3[0] {
		return top3
	}

	top3[0] = val
	sort.Sort(top3)
	log.Print(top3)
	return top3
}

func Part2() {
	lines := readInput()

	elfMaxes := make(TopThree, 0, 3)
	var elfTotal int

	for _, line := range lines {
		if line == "" {
			elfMaxes = tryInsert(elfMaxes, elfTotal)
			elfTotal = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		elfTotal += val
	}

	sum := 0
	for _, val := range elfMaxes {
		sum += val
	}
	log.Print(sum)
}
