package day09

import (
	"fmt"
	"io"
	"math"
	"slices"
	"sort"
	"time"

	"github.com/minyiky/advent-of-code/2025/helpers/point"
)

type segment struct {
	start, end int
}

func buildHorizontalData(polyPoints map[point.Point]struct{}) map[int][]segment {
	boundaryByRow := make(map[int][]int)
	for p := range polyPoints {
		boundaryByRow[p.Y] = append(boundaryByRow[p.Y], p.X)
	}

	insideRangesByRow := make(map[int][]segment)
	for y, xs := range boundaryByRow {
		sort.Ints(xs)

		var segs []segment
		i := 0
		for i < len(xs) {
			start := xs[i]
			for i+1 < len(xs) && xs[i+1] == xs[i]+1 {
				i++
			}
			segs = append(segs, segment{start: start, end: xs[i]})
			i++
		}

		var insideRanges []segment
		for i := 0; i+1 < len(segs); i += 2 {
			insideRanges = append(insideRanges, segment{start: segs[i].start, end: segs[i+1].end})
		}
		insideRangesByRow[y] = insideRanges
	}

	return insideRangesByRow
}

func isRangeInside(y, xStart, xEnd int, rowRanges map[int][]segment) bool {
	ranges := rowRanges[y]
	idx := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].end >= xStart
	})
	return idx < len(ranges) && ranges[idx].start <= xStart && xEnd <= ranges[idx].end
}

func buildVerticalData(rowRanges map[int][]segment, candidateXs, criticalYs []int) map[int][]segment {
	colRanges := make(map[int][]segment, len(candidateXs))

	for _, x := range candidateXs {
		var segs []segment
		segStart, inSeg := 0, false

		for i, y := range criticalYs {
			inside := isRangeInside(y, x, x, rowRanges)

			if inside && !inSeg {
				inSeg, segStart = true, y
			} else if !inside && inSeg {
				segs = append(segs, segment{start: segStart, end: criticalYs[i-1]})
				inSeg = false
			}
		}

		if inSeg {
			segs = append(segs, segment{start: segStart, end: criticalYs[len(criticalYs)-1]})
		}

		colRanges[x] = segs
	}
	return colRanges
}

func isVerticalRangeInside(x, yStart, yEnd int, colRanges map[int][]segment) bool {
	segs := colRanges[x]
	idx := sort.Search(len(segs), func(i int) bool {
		return segs[i].end >= yStart
	})
	return idx < len(segs) && segs[idx].start <= yStart && yEnd <= segs[idx].end
}

func Part2Val(lines []string) (int, error) {
	minX, minY := math.MaxInt, math.MaxInt

	points := make([]point.Point, 0, len(lines))
	for _, line := range lines {
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			return 0, err
		}
		points = append(points, point.Point{X: x, Y: y})
		minX, minY = min(minX, x), min(minY, y)
	}

	xsSet := make(map[int]struct{}, len(points))
	ysSet := make(map[int]struct{}, len(points))
	for i, p := range points {
		p.X -= minX
		p.Y -= minY
		points[i] = p
		xsSet[p.X] = struct{}{}
		ysSet[p.Y] = struct{}{}
	}

	polyPoints := make(map[point.Point]struct{})
	for i := 0; i < len(points); i++ {
		p1, p2 := points[i], points[(i+1)%len(points)]
		dx, dy := sign(p2.X-p1.X), sign(p2.Y-p1.Y)
		for p := p1; p != p2; p = p.Add(point.Point{X: dx, Y: dy}) {
			polyPoints[p] = struct{}{}
		}
		polyPoints[p2] = struct{}{}
	}

	rowRanges := buildHorizontalData(polyPoints)

	candidateXs := sortedKeys(xsSet)
	criticalYs := sortedKeys(ysSet)
	colRanges := buildVerticalData(rowRanges, candidateXs, criticalYs)

	slices.SortFunc(points, func(a, b point.Point) int {
		return b.Magnitude() - a.Magnitude()
	})

	currentMax := 0
	for i, p := range points {
		if p.Magnitude() <= currentMax {
			break
		}

		for _, otherP := range points[i+1:] {
			minX, maxX := min(p.X, otherP.X), max(p.X, otherP.X)
			minY, maxY := min(p.Y, otherP.Y), max(p.Y, otherP.Y)

			size := (maxX - minX + 1) * (maxY - minY + 1)
			if size <= currentMax {
				continue
			}

			if !isRangeInside(minY, minX, maxX, rowRanges) ||
				!isRangeInside(maxY, minX, maxX, rowRanges) ||
				!isVerticalRangeInside(minX, minY, maxY, colRanges) ||
				!isVerticalRangeInside(maxX, minY, maxY, colRanges) {
				continue
			}

			currentMax = size
		}
	}

	return currentMax, nil
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func sortedKeys(m map[int]struct{}) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func Part2(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part2Val(lines)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "The value found was: %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(time.Since(start))/1e6)
	return nil
}
