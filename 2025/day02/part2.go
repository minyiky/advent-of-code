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

// Part2ValWithViz runs Part2 with visualization support
func Part2ValWithViz(lines []string, vizCallback func(data map[string]interface{}), speed float64) (int, error) {
	var value int

	valueSet := make(map[int]struct{})
	rngs := strings.Split(lines[0], ",")

	for rangeIdx, line := range rngs {
		var lower, upper int
		if _, err := fmt.Sscanf(line, "%d-%d", &lower, &upper); err != nil {
			return 0, err
		}

		lowerLen := len(strconv.Itoa(lower))
		upperLen := len(strconv.Itoa(upper))

		// Collect ALL patterns for this specific range (all repetition counts)
		panels := make(map[string][]string)
		rangeValueSet := make(map[int]int) // value -> repetition count

		for l := 2; l <= upperLen; l++ {
			start, _ := func() (int, error) {
				if lowerLen%l != 0 {
					return int(math.Pow10(lowerLen / l)), nil
				}
				return strconv.Atoi(strconv.Itoa(lower)[:lowerLen/l])
			}()
			end, _ := strconv.Atoi(strconv.Itoa(upper)[:int(math.Ceil(float64(upperLen)/float64(l)))])

			// Process values in this repetition count
			if start > 0 {
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

					// Only add if new (deduplication within this range display)
					if _, exists := rangeValueSet[val]; !exists {
						rangeValueSet[val] = l

						// Add to global set
						valueSet[val] = struct{}{}
					}
				}
			}
		}

		// Build panels from collected values for this range
		for val, repCount := range rangeValueSet {
			// Determine which power range panel this belongs to
			valLen := len(strconv.Itoa(val))
			powerRangeStart := int(math.Pow10(valLen - 1))
			powerRangeEnd := int(math.Pow10(valLen)) - 1
			powerRangeLabel := fmt.Sprintf("%d-%d", powerRangeStart, powerRangeEnd)

			valStr := fmt.Sprintf("%d (×%d)", val, repCount)
			panels[powerRangeLabel] = append(panels[powerRangeLabel], valStr)
		}

		// Send ONE update showing all patterns found in THIS range
		if vizCallback != nil {
			// Create a deep copy of panels for safe channel transmission
			panelsCopy := make(map[string]interface{})
			for k, v := range panels {
				valuesCopy := make([]interface{}, len(v))
				for i, val := range v {
					valuesCopy[i] = val
				}
				panelsCopy[k] = valuesCopy
			}

			vizCallback(map[string]interface{}{
				"currentRange": line,
				"rangeIndex":   rangeIdx,
				"totalRanges":  len(rngs),
				"panels":       panelsCopy,
				"currentSum":   calculateSum(valueSet),
			})
		}
	}

	for v := range valueSet {
		value += v
	}
	return value, nil
}

// Helper function to calculate sum from value set
func calculateSum(valueSet map[int]struct{}) int {
	sum := 0
	for v := range valueSet {
		sum += v
	}
	return sum
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
