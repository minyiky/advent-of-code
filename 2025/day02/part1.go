package day02

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	rngs := strings.Split(lines[0], ",")

	for _, line := range rngs {
		var lower, upper int
		if _, err := fmt.Sscanf(line, "%d-%d", &lower, &upper); err != nil {
			return 0, err
		}

		lowerLen := len(strconv.Itoa(lower))
		upperLen := len(strconv.Itoa(upper))

		start, _ := func() (int, error) {
			if lowerLen%2 != 0 {
				return int(math.Pow10(lowerLen / 2)), nil
			}
			return strconv.Atoi(strconv.Itoa(lower)[:lowerLen/2])
		}()
		end, _ := strconv.Atoi(strconv.Itoa(upper)[:int(math.Ceil(float64(upperLen)/2))])

		for i := start; i <= end; i++ {
			iStr := strconv.Itoa(i)
			val, _ := strconv.Atoi(iStr + iStr)

			if val > upper {
				break
			}

			if val < lower {
				continue
			}

			value += val
		}
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
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
