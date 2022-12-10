package day08

import (
	"fmt"
	"io"
	"math"
)

func getNext(height rune, view []rune) int {
	if len(view) == 0 {
		return 0
	}
	var num int
	for _, tree := range view {
		num++
		if tree >= height {
			break
		}
	}
	return num
}

func Part2Val(lines []string) int {
	var value = math.MinInt

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		for _, char := range line {
			grid[i] = append(grid[i], char-48)
		}
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			height := grid[i][j]
			upSlice := []rune{}
			for k := j - 1; k > -1; k-- {
				upSlice = append(upSlice, grid[i][k])
			}
			up := getNext(height, upSlice)

			downSlice := grid[i][j+1:]
			down := getNext(height, downSlice)

			rightSlice := []rune{}
			for k := i + 1; k < len(grid); k++ {
				rightSlice = append(rightSlice, grid[k][j])
			}
			right := getNext(height, rightSlice)

			leftSlice := []rune{}
			for k := i - 1; k > -1; k-- {
				leftSlice = append(leftSlice, grid[k][j])
			}
			left := getNext(height, leftSlice)

			view := up * down * right * left
			if view > value {
				value = view
			}
		}
	}

	return value
}

func Part2(w io.Writer, lines []string) {
	value := Part2Val(lines)
	fmt.Fprintf(w, "The elves found that the best view had a total score of  %d\n", value)
}
