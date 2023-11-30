package day01

import (
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Integer | constraints.Float | constraints.Complex](vals []T) T {
	var sum T
	for _, val := range vals {
		sum += val
	}
	return sum
}

func Part2Val(lines []int) (int, error) {
	var numDeeper int
	for i := 3; i < len(lines); i++ {
		if sum(lines[i-3:i]) < sum(lines[i-2:i+1]) {
			numDeeper++
		}
	}
	return numDeeper, nil
}

func Part2(w io.Writer, lines []int) error {
	val, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Using a sliding window, the depth measurement increased %d times\n", val)
	return nil
}
