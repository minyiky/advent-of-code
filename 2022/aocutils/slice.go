package aocutils

import "golang.org/x/exp/constraints"

func ReverseSlice[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func SliceContains[T comparable](slice []T, item T) (int, bool) {
	for i, val := range slice {
		if val == item {
			return i, true
		}
	}

	return 0, false
}

func SliceMax[T constraints.Ordered](slice []T) T {
	max := slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}
