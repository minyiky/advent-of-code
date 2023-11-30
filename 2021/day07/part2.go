package day07

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
)

func Part2Val(lines []string) (int, error) {
	value := math.MaxInt

	vals := strings.Split(lines[0], ",")

	nums := make([]int, len(vals))

	max := math.MinInt
	min := math.MaxInt

	for i, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			return 0, nil
		}
		max = maths.Max(num, max)
		if num < min {
			min = num
		}
		nums[i] = num
	}

	value = binarySearch(min, max, func(modifier int) int {
		val := 0
		for _, num := range nums {
			val += triangular(maths.Abs(num - modifier))
		}
		return val
	})

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
