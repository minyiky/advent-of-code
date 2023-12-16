package day16

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

var (
	up    = point.NewPoint2D(0, -1)
	down  = point.NewPoint2D(0, 1)
	right = point.NewPoint2D(1, 0)
	left  = point.NewPoint2D(-1, 0)
)

func reflect(dir point.Point2D, char rune) []point.Point2D {
	switch {
	case (dir == up && char == '/') ||
		(dir == down && char == '\\'):
		return []point.Point2D{right}
	case dir == up && char == '\\' ||
		(dir == down && char == '/'):
		return []point.Point2D{left}
	case (dir == up || dir == down) && char == '-':
		return []point.Point2D{left, right}
	case (dir == right && char == '/') ||
		(dir == left && char == '\\'):
		return []point.Point2D{up}
	case (dir == right && char == '\\') ||
		(dir == left && char == '/'):
		return []point.Point2D{down}
	case (dir == right || dir == left) && char == '|':
		return []point.Point2D{up, down}
	}
	return []point.Point2D{dir}
}

func move(pos, dir point.Point2D, lines []string, cache map[point.Point2D]map[point.Point2D]bool) {
	// fmt.Println(pos, dir, string(lines[pos.Y()][pos.X()]))
	if posCache, ok := cache[pos]; ok {
		if _, ok2 := posCache[dir]; ok2 {
			return
		}
	}

	if _, ok := cache[pos]; !ok {
		cache[pos] = make(map[point.Point2D]bool)
	}

	cache[pos][dir] = true

	for _, d := range reflect(dir, rune(lines[pos.Y()][pos.X()])) {
		newPos := point.Add(pos, d)
		if newPos.X() >= 0 && newPos.X() < len(lines[0]) &&
			newPos.Y() >= 0 && newPos.Y() < len(lines) {
			move(newPos, d, lines, cache)
		}
	}
}

func Part1Val(lines []string) (int, error) {
	var value int

	cache := make(map[point.Point2D]map[point.Point2D]bool)

	move(point.NewPoint2D(0, 0), right, lines, cache)

	value = len(cache)

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
