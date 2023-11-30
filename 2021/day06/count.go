package day06

import (
	"strconv"
	"strings"
)

func count(lines []string, days int) (int, error) {
	var value int

	fish := make(map[string]int)

	for i := 0; i < 9; i++ {
		fish[strconv.Itoa(i)] = 0
	}

	inputs := strings.Split(lines[0], ",")

	for _, input := range inputs {
		fish[input]++
	}

	for i := 0; i < days; i++ {
		newFish := make(map[string]int)
		for i := 8; i >= 0; i-- {
			if i == 0 {
				newFish["8"] = fish["0"]
				newFish["6"] += fish["0"]
				continue
			}
			newFish[strconv.Itoa(i-1)] = fish[strconv.Itoa(i)]
		}
		fish = newFish
	}

	for _, v := range fish {
		value += v
	}
	return value, nil
}
