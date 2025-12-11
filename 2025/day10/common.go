package day10

import "sort"

// key takes a slice of a slice of ints and outputs a string with an ordered list of ints e.g.:
//
//	[[1, 2] [3]] -> "123"
//	[[5], [1, 2]] -> "125"
//	[[1, 2], [3, 1]] -> "1123"
func key(arr [][]int) string {
	var vals []int
	for _, row := range arr {
		vals = append(vals, row...)
	}
	sort.Ints(vals)
	var runes []rune
	for _, v := range vals {
		runes = append(runes, rune(v+'0'))
	}
	return string(runes)
}

func extendKey(k string, arr []int) string {
	newKey := key([][]int{arr})
	combined := k + newKey
	runes := []rune(combined)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
