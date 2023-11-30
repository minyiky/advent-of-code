package day07

import (
	"math"
)

func linearSearch(min, max int, eval func(int) int) int {
	value := math.MaxInt
	for i := min; i <= max; i++ {
		val := eval(i)
		if val < value {
			value = val
		}
	}
	return value
}

func binarySearch(min, max int, eval func(int) int) int {
	midPoint := (max + min) / 2

	if min == max {
		return eval(min)
	}

	upperPoint := (max + midPoint) / 2
	upperVal := eval(upperPoint)

	lowerPoint := (min + midPoint) / 2
	lowerVal := eval(lowerPoint)

	if upperVal < lowerVal {
		return binarySearch(midPoint, max, eval)
	} else {
		return binarySearch(min, midPoint, eval)
	}
}
