package day15

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

func search(pos, end point.Point2D, base, costs map[point.Point2D]int, best int, num *int) int {
	if pos == end {
		*num += 1
		return costs[pos]
	}

	do := func(d point.Point2D) {
		newPos := point.Add(pos, d)
		if _, ok := base[newPos]; !ok {
			return
		}
		newCost := costs[pos] + base[newPos]
		if newCost < costs[newPos] && newCost < best {
			costs[newPos] = costs[pos] + base[newPos]
			if val := search(newPos, end, base, costs, best, num); val < best {
				best = min(val, best)
			}
		}
	}

	// directions := []point.Point2D{}

	// if closest(pos, end) == point.NewPoint2D(1, 0) {
	// 	directions = append(directions, point.NewPoint2D(1, 0), point.NewPoint2D(0, 1))
	// } else {
	// 	directions = append(directions, point.NewPoint2D(0, 1), point.NewPoint2D(1, 0))
	// }

	// directions = append(directions, backwards...)

	for _, d := range allDirections {
		do(d)
	}

	return best
}

func searchPriority(pos, end point.Point2D, base, costs map[point.Point2D]int, best int, num *int) int {
	if pos == end {
		// fmt.Println(costs[pos])
		*num += 1
		return costs[pos]
	}

	do := func(d point.Point2D) {
		newPos := point.Add(pos, d)
		newCost := costs[pos] + base[newPos]
		if newCost < costs[newPos] && newCost < best {
			costs[newPos] = costs[pos] + base[newPos]
			if val := searchPriority(newPos, end, base, costs, best, num); val < best {
				best = min(val, best)
			}
		}
	}

	type costList struct {
		p point.Point2D
		c int
	}

	fList := make([]costList, 0, 4)

	for _, d := range allDirections[:2] {
		newPos := point.Add(pos, d)
		if _, ok := base[newPos]; !ok {
			continue
		}
		fList = append(fList, costList{p: d, c: base[newPos]})
	}

	slices.SortFunc(fList, func(a costList, b costList) int {
		return a.c - b.c
	})

	bList := make([]costList, 0, 2)

	for _, d := range allDirections[:2] {
		newPos := point.Add(pos, d)
		if _, ok := base[newPos]; !ok {
			continue
		}
		bList = append(bList, costList{p: d, c: base[newPos]})
	}

	slices.SortFunc(bList, func(a costList, b costList) int {
		return a.c - b.c
	})

	list := append(fList, bList...)

	// directions := []point.Point2D{}

	// if closest(pos, end) == point.NewPoint2D(1, 0) {
	// 	directions = append(directions, point.NewPoint2D(1, 0), point.NewPoint2D(0, 1))
	// } else {
	// 	directions = append(directions, point.NewPoint2D(0, 1), point.NewPoint2D(1, 0))
	// }

	// directions = append(directions, backwards...)

	for _, d := range list {
		do(d.p)
	}

	return best
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2021 day 15 --\n")
	if err := Part1(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
