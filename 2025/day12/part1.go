package day12

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part1Val(lines []string) (int, error) {
	var value int

	for _, line := range lines {
		size, blockSet, found := strings.Cut(line, ":")
		if !found {
			continue
		}

		var w, h int
		_, err := fmt.Sscanf(size, "%dx%d", &w, &h)
		if err != nil {
			continue
		}

		w = (w / 3)
		h = (h / 3)

		blocks := strings.Fields(blockSet)
		area := w * h
		blockArea := 0
		for _, block := range blocks {
			numBlocks, err := strconv.Atoi(block)
			if err != nil {
				continue
			}
			blockArea += numBlocks
		}

		if area >= blockArea {
			value += 1
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
