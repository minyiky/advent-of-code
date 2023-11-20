package day20

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

func Part2Val(lines []string) (int, error) {
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
		num := Number{value: v * 811589153, key: i}
		list[i] = num
		keys[num] = i
		if v == 0 {
			Zero = Number{value: v, key: i}
		}
	}

	newList := container.CopySlice(list)

	for x := 0; x < 10; x++ {
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

			var sliceStart, sliceEnd int
			if start < end {
				sliceStart = start
				sliceEnd = end + 1
			} else {
				sliceStart = end
				sliceEnd = start + 1
			}
			for i := sliceStart; i < sliceEnd; i++ {
				keys[newList[i]] = i
			}
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		value += newList[(keys[Zero]+v)%length].value
	}
	return value, nil
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "Having remembered the encryption key you found that the real coordinates are %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
