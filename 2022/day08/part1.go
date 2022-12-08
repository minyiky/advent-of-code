package day08

import (
	"log"
)

func Part1Val(lines []string) (int, error) {
	var value int

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		for _, char := range line {
			grid[i] = append(grid[i], char)
		}
	}

	seen := make(map[[2]int]bool)
	top := make([]rune, len(grid))
	bottom := make([]rune, len(grid))
	left := make([]rune, len(grid[0]))
	right := make([]rune, len(grid[0]))
	height := len(grid)
	for i := 0; i < len(grid); i++ {
		width := len(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > top[i] {
				seen[[2]int{i, j}] = true
				top[i] = grid[i][j]
			}
			if grid[i][width-j-1] > bottom[i] {
				seen[[2]int{i, width - j - 1}] = true
				bottom[i] = grid[i][width-j-1]
			}
			if grid[i][j] > left[j] {
				seen[[2]int{i, j}] = true
				left[j] = grid[i][j]
			}
			if grid[height-i-1][j] > right[j] {
				seen[[2]int{height - i - 1, j}] = true
				right[j] = grid[height-i-1][j]
			}
		}
	}

	for k, _ := range seen {
		_ = k
		value++
	}

	return value, nil
}

func Part1(lines []string) error {
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	log.Printf("From outside the new patch of trees, %d trees could be seen", value)
	return nil
}
