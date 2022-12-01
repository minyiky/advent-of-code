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

func Part2() {
	elfMaxes := make(TopThree, 0, len(lines))
	var elfTotal int

	for _, line := range lines {
		if line == "" {
			elfMaxes = append(elfMaxes, elfTotal)
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

	sort.Sort(elfMaxes)

	sum := 0
	for _, val := range elfMaxes[len(elfMaxes)-3:] {
		sum += val
	}
	log.Printf("The top three elves carried a total of %d calories\n", sum)
}
