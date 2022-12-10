package day08

import (
	"fmt"
	"io"
)

func Part1Val(lines []string) int {
	var value int

	// initialise the first axis of the array
	grid := make([][]rune, len(lines))

	// Add all characters to the 2d Grid (slow due to lots of appends)
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
				if _, ok := seen[[2]int{i, j}]; !ok {
					value++
					seen[[2]int{i, j}] = true
				}
				top[i] = grid[i][j]
			}
			if grid[i][width-j-1] > bottom[i] {
				if _, ok := seen[[2]int{i, width - j - 1}]; !ok {
					value++
					seen[[2]int{i, width - j - 1}] = true
				}
				bottom[i] = grid[i][width-j-1]
			}
			if grid[i][j] > left[j] {
				if _, ok := seen[[2]int{i, j}]; !ok {
					value++
					seen[[2]int{i, j}] = true
				}
				left[j] = grid[i][j]
			}
			if grid[height-i-1][j] > right[j] {
				if _, ok := seen[[2]int{height - i - 1, j}]; !ok {
					value++
					seen[[2]int{height - i - 1, j}] = true
				}
				right[j] = grid[height-i-1][j]
			}
		}
	}

	return value
}

func Part1(w io.Writer, lines []string) {
	value := Part1Val(lines)
	fmt.Fprintf(w, "From outside the new patch of trees, %d trees could be seen\n", value)
}
