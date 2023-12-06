package day06

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"time"
)

func Part2Val(lines []string) (int, error) {
	value := 1

	tStrs := rNum.FindAllString(lines[0], -1)
	dStrs := rNum.FindAllString(lines[1], -1)

	convert := func(strs []string) float64 {
		str := ""
		for _, s := range strs {
			str += s
		}
		num, _ := strconv.Atoi(str)
		return float64(num)
	}

	t := convert(tStrs)
	d := convert(dStrs)

	low, high := quad(d, t)
	lowI := math.Ceil(low)
	if test(int(d), int(t), int(lowI)) {
		lowI++
	}
	highI := math.Floor(high)
	if test(int(d), int(t), int(highI)) {
		highI--
	}
	value *= int(highI - lowI + 1)

	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "There were %d ways to win the long race\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
