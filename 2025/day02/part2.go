package day02

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func Part2Val(lines []string) (int, error) {
	var value int

	valueSet := make(map[int]struct{})
	rngs := strings.Split(lines[0], ",")

	for _, line := range rngs {
		var lower, upper int
		if _, err := fmt.Sscanf(line, "%d-%d", &lower, &upper); err != nil {
			return 0, err
		}

		lowerLen := len(strconv.Itoa(lower))
		upperLen := len(strconv.Itoa(upper))

		for l := 2; l <= upperLen; l++ {
			start, _ := func() (int, error) {
				if lowerLen%l != 0 {
					return int(math.Pow10(lowerLen / l)), nil
				}
				return strconv.Atoi(strconv.Itoa(lower)[:lowerLen/l])
			}()
			end, _ := strconv.Atoi(strconv.Itoa(upper)[:int(math.Ceil(float64(upperLen)/float64(l)))])

			for i := start; i <= end; i++ {
				iStr := strconv.Itoa(i)
				vStr := ""
				for range l {
					vStr += iStr
				}
				val, _ := strconv.Atoi(vStr)

				if val > upper {
					break
				}

				if val < lower {
					continue
				}

				valueSet[val] = struct{}{}
			}
		}
	}

	for v := range valueSet {
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
