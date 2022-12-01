package dayone

import (
	"log"
	"os"
	"strconv"
)

func Part1() {
	lines := readInput()

	var elfTotal, elfMax int

	for _, line := range lines {
		if line == "" {
			if elfTotal > elfMax {
				elfMax = elfTotal
			}
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

	log.Println(elfMax)
}
