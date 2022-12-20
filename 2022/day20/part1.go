package day20

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func insert[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func Part1Val(lines []string) (int, error) {
	var value int

	length := len(lines)
	fmt.Println(length)

	list := make([]Number, length)
	keys := make(map[Number]int)
	Zero := Number{}

	for i, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		num := Number{value: v, key: i}
		list[i] = num
		keys[num] = i
		if v == 0 {
			Zero = Number{value: v, key: i}
		}
	}

	newList := aocutils.CopySlice(list)

	for _, num := range list {
		start := keys[num]
		diff := num.value
		end := start + diff

		end %= (length - 1)
		if end < 0 {
			end += length - 1
		}

		if start == end {
			continue
		}

		if start < length-1 {
			newList = append(newList[:start], newList[start+1:]...)
		} else {
			newList = newList[:start]
		}

		newList = insert(newList, end, num)

		for i := 0; i < length; i++ {
			keys[newList[i]] = i
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		value += newList[(keys[Zero]+v)%length].value
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
	fmt.Fprintf(w, "Having run the algorithm you find that the coordinates are %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
