package dayone

import (
	"log"
	"os"
	"path/filepath"

	"github.com/minyiky/advent-of-code/2022/utils"
)

func readInput() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	path := filepath.Join(dir, "1/input.txt")
	lines, err := utils.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return lines
}
