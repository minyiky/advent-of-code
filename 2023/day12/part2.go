package day12

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Key struct {
	pos, cog, length int
}

func dp(cogPart string, cogs []int, pos, cog, length int, cache map[Key]int) int {
	if pos == len(cogPart) {
		if (length == 0 && cog == len(cogs)) ||
			(length == cogs[len(cogs)-1] && cog == len(cogs)-1) {
			return 1
		} else {
			return 0
		}
	}

	key := Key{
		pos,
		cog,
		length,
	}

	if val, ok := cache[key]; ok {
		return val
	}

	var val int

	for _, char := range []rune{'.', '#'} {
		if !(cogPart[pos] == '?' || rune(cogPart[pos]) == char) {
			continue
		}
		if char == '.' && length == 0 {
			val += dp(cogPart, cogs, pos+1, cog, 0, cache)
		} else if char == '.' && length == cogs[cog] {
			val += dp(cogPart, cogs, pos+1, cog+1, 0, cache)
		} else if char == '#' && cog < len(cogs) && length < cogs[cog] {
			val += dp(cogPart, cogs, pos+1, cog, length+1, cache)

		}
	}

	cache[key] = val
	return val
}

func Part2Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {

		cogPart, numPart, _ := strings.Cut(line, " ")

		cogPart = strings.Join([]string{
			cogPart,
			cogPart,
			cogPart,
			cogPart,
			cogPart,
		}, "?")

		numPart = strings.Join([]string{
			numPart,
			numPart,
			numPart,
			numPart,
			numPart,
		}, ",")

		numsStr := rNum.FindAllString(numPart, -1)
		nums := make([]int, len(numsStr))
		for j, numStr := range numsStr {
			nums[j], _ = strconv.Atoi(numStr)
		}

		cache := make(map[Key]int)
		v := dp(cogPart, nums, 0, 0, 0, cache)
		value += v
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
