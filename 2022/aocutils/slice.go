package aocutils

import "golang.org/x/exp/constraints"

// ReverseSlice changes the order of elements in a slice in place
func ReverseSlice[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// SliceContains returns whether an item exists in a slice,
// and the first location it is found at
func SliceContains[T comparable](slice []T, item T) (int, bool) {
	for i, val := range slice {
		if val == item {
			return i, true
		}
	}
	return 0, false
}

// SliceMax returns the maximum value found in a slice
func SliceMax[T constraints.Ordered](slice []T) T {
	max := slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}
