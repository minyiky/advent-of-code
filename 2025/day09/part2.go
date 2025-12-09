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

type polyData struct {
	insideRangesByRow map[int][]segment
	polyPoints        map[point.Point]struct{}
}

func buildPolyData(polyPoints map[point.Point]struct{}) polyData {
	boundaryByRow := make(map[int][]int)
	for p := range polyPoints {
		boundaryByRow[p.Y] = append(boundaryByRow[p.Y], p.X)
	}

	boundarySegmentsByRow := make(map[int][]segment)
	for y, xs := range boundaryByRow {
		sort.Ints(xs)
		var segs []segment
		i := 0
		for i < len(xs) {
			start := xs[i]
			end := start
			for i+1 < len(xs) && xs[i+1] == xs[i]+1 {
				i++
				end = xs[i]
			}
			segs = append(segs, segment{start: start, end: end})
			i++
		}
		boundarySegmentsByRow[y] = segs
	}

	insideRangesByRow := make(map[int][]segment)
	for y, segs := range boundarySegmentsByRow {
		var insideRanges []segment
		inside := false
		for i := 0; i < len(segs); i++ {
			if !inside {
				insideRanges = append(insideRanges, segment{start: segs[i].start, end: segs[i].end})
				inside = true
			} else {
				insideRanges[len(insideRanges)-1].end = segs[i].end
				inside = false
			}
		}
		insideRangesByRow[y] = insideRanges
	}

	return polyData{
		insideRangesByRow: insideRangesByRow,
		polyPoints:        polyPoints,
	}
}

func isRangeInside(y, xStart, xEnd int, pd *polyData) bool {
	ranges := pd.insideRangesByRow[y]
	if len(ranges) == 0 {
		return false
	}

	idx := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].end >= xStart
	})

	if idx >= len(ranges) {
		return false
	}

	return ranges[idx].start <= xStart && xEnd <= ranges[idx].end
}

func buildVerticalData(pd *polyData, candidateXs []int, criticalYs []int) map[int][]segment {
	colRanges := make(map[int][]segment, len(candidateXs))

	for _, x := range candidateXs {
		var segs []segment
		inSeg := false
		segStart := 0

		for i := 0; i < len(criticalYs)-1; i++ {
			y := criticalYs[i]
			nextY := criticalYs[i+1]

			inside := isRangeInside(y, x, x, pd)
			if inside {
				if !inSeg {
					inSeg = true
					segStart = y
				}
			} else {
				if inSeg {
					segs = append(segs, segment{start: segStart, end: y - 1})
					inSeg = false
				}
			}

			midY := (y + nextY) / 2
			if midY > y && midY < nextY {
				midInside := isRangeInside(midY, x, x, pd)
				if midInside {
					if !inSeg {
						inSeg = true
						segStart = y + 1
					}
				} else {
					if inSeg {
						segs = append(segs, segment{start: segStart, end: y})
						inSeg = false
					}
				}
			}
		}

		lastY := criticalYs[len(criticalYs)-1]
		if isRangeInside(lastY, x, x, pd) {
			if !inSeg {
				segs = append(segs, segment{start: lastY, end: lastY})
			} else {
				segs = append(segs, segment{start: segStart, end: lastY})
			}
		} else if inSeg {
			segs = append(segs, segment{start: segStart, end: lastY - 1})
		}

		colRanges[x] = segs
	}

	return colRanges
}

func isVerticalRangeInside(x, yStart, yEnd int, colRanges map[int][]segment) bool {
	segs := colRanges[x]
	if len(segs) == 0 {
		return false
	}

	idx := sort.Search(len(segs), func(i int) bool {
		return segs[i].end >= yStart
	})
	if idx >= len(segs) {
		return false
	}

	return segs[idx].start <= yStart && yEnd <= segs[idx].end
}

func Part2Val(lines []string) (int, error) {
	minX, minY := math.MaxInt, math.MaxInt
	maxY := 0

	points := make([]point.Point, 0, len(lines))

	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return 0, err
		}
		points = append(points, point.Point{X: x, Y: y})
		minX = min(minX, x)
		minY = min(minY, y)
	}

	for i, p := range points {
		p.X -= minX
		p.Y -= minY
		points[i] = p
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	polyPoints := make(map[point.Point]struct{})

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]

		dx := p2.X - p1.X
		dy := p2.Y - p1.Y

		steps := int(max(math.Abs(float64(dx)), math.Abs(float64(dy))))

		if steps > 0 {
			dx /= int(math.Abs(float64(steps)))
			dy /= int(math.Abs(float64(steps)))
		}

		for step := 0; step <= steps; step++ {
			x := p1.X + dx*step
			y := p1.Y + dy*step
			polyPoints[point.Point{X: x, Y: y}] = struct{}{}
		}
	}

	pd := buildPolyData(polyPoints)

	xsSet := make(map[int]struct{}, len(points))
	ysSet := make(map[int]struct{}, len(points))
	for _, p := range points {
		xsSet[p.X] = struct{}{}
		ysSet[p.Y] = struct{}{}
	}
	candidateXs := make([]int, 0, len(xsSet))
	for x := range xsSet {
		candidateXs = append(candidateXs, x)
	}
	sort.Ints(candidateXs)

	criticalYs := make([]int, 0, len(ysSet))
	for y := range ysSet {
		criticalYs = append(criticalYs, y)
	}
	sort.Ints(criticalYs)

	colRanges := buildVerticalData(&pd, candidateXs, criticalYs)

	slices.SortFunc(points, func(a, b point.Point) int {
		return b.Magnitude() - a.Magnitude()
	})

	currentMax := 0

	for i, p := range points {
		if p.Magnitude() <= currentMax {
			break
		}

	pointLoop:
		for _, otherP := range points[i+1:] {
			if p.X <= otherP.X || p.Y == otherP.Y {
				continue pointLoop
			}

			minPX := min(p.X, otherP.X)
			maxPX := max(p.X, otherP.X)
			minPY := min(p.Y, otherP.Y)
			maxPY := max(p.Y, otherP.Y)

			width := maxPX - minPX + 1
			height := maxPY - minPY + 1
			size := width * height

			if size <= currentMax {
				continue pointLoop
			}

			if !isRangeInside(minPY, minPX, maxPX, &pd) {
				continue pointLoop
			}
			if !isRangeInside(maxPY, minPX, maxPX, &pd) {
				continue pointLoop
			}

			if !isVerticalRangeInside(minPX, minPY, maxPY, colRanges) {
				continue pointLoop
			}
			if !isVerticalRangeInside(maxPX, minPY, maxPY, colRanges) {
				continue pointLoop
			}

			currentMax = size
		}
	}

	return currentMax, nil
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
