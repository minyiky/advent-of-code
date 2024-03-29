package day18

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

type Limits struct {
	minX, minY, minZ, maxX, maxY, maxZ int
}

func (l Limits) Outside(p point.Point3D) bool {
	return p.X() <= l.minX || p.X() >= l.maxX || p.Y() <= l.minY || p.Y() >= l.maxY || p.Z() <= l.minZ || p.Z() >= l.maxZ
}

func WayOut(pos point.Point3D, blocks, seen, outside map[point.Point3D]bool, limits Limits) bool {
	if blocks[pos] {
		return false
	}

	if outside[pos] || limits.Outside(pos) {
		return true
	}

	for _, direction := range directions {
		newPoint := point.Add(pos, direction)
		if !seen[newPoint] {
			seen[newPoint] = true
			if WayOut(newPoint, blocks, seen, outside, limits) {
				return true
			}
		}
	}
	return false
}

func Part2Val(lines []string) (int, error) {
	var value int
	var limits Limits

	grid := make(map[point.Point3D]bool)
	isOutside := make(map[point.Point3D]bool)

	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		grid[point.NewPoint3D(x, y, z)] = true
		if x < limits.minX {
			limits.minX = x
		}
		if x > limits.maxX {
			limits.maxX = x
		}
		if y < limits.minY {
			limits.minY = y
		}
		if y > limits.maxY {
			limits.maxY = y
		}
		if z < limits.minZ {
			limits.minZ = z
		}
		if z > limits.maxZ {
			limits.maxZ = z
		}
	}

	for pos := range grid {
		for _, direction := range directions {
			newPoint := point.Add(pos, direction)
			if _, ok := grid[newPoint]; !ok {
				ok, tested := isOutside[newPoint]
				if !tested {
					testedMap := make(map[point.Point3D]bool)
					outside := WayOut(newPoint, grid, testedMap, isOutside, limits)
					if outside {
						value++
						for seen := range testedMap {
							isOutside[seen] = true
						}
					} else {
						for seen := range testedMap {
							isOutside[seen] = false
						}
					}
				} else if ok {
					value++
				}
			}
		}
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
	fmt.Fprintf(w, "Only counting the faces with a path to the surface, %d were uncovered\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
