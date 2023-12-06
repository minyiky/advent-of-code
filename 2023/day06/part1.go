package day06

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"time"
)

func quad(d, t float64) (float64, float64) {
	low := (t - math.Sqrt(t*t-4*d)) / 2
	high := (t + math.Sqrt(t*t-4*d)) / 2
	return low, high
}

func test(d, t, x int) bool {
	return x*(t-x) == d
}

func Part1Val(lines []string) (int, error) {
	value := 1

	tStrs := rNum.FindAllString(lines[0], -1)
	dStrs := rNum.FindAllString(lines[1], -1)

	convert := func(strs []string) []float64 {
		nums := make([]float64, len(strs))
		for i := range strs {
			num, _ := strconv.Atoi(strs[i])
			nums[i] = float64(num)
		}
		return nums
	}

	ts := convert(tStrs)
	ds := convert(dStrs)

	for i := range ts {
		low, high := quad(ds[i], ts[i])
		lowI := math.Ceil(low)
		if test(int(ds[i]), int(ts[i]), int(lowI)) {
			lowI++
		}
		highI := math.Floor(high)
		if test(int(ds[i]), int(ts[i]), int(highI)) {
			highI--
		}
		value *= int(highI - lowI + 1)
	}

	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "There were %d ways to win the different races\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
