package day14

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

func getCoords(point string) (int, int, error) {
	points := strings.Split(point, ",")
	if len(points) != 2 {
		return 0, 0, errors.New("wrong number of point sections")
	}

	x, err := strconv.Atoi(points[0])
	if err != nil {
		return 0, 0, fmt.Errorf("value not an int: %s", points[0])
	}
	y, err := strconv.Atoi(points[1])
	if err != nil {
		return 0, 0, fmt.Errorf("value not an int: %s", points[1])
	}

	return x, y, nil
}

func createMap(lines []string) (map[point.Point2D]bool, int, error) {

	blocked := make(map[point.Point2D]bool)
	var yBig int

	for _, line := range lines {
		paths := strings.Split(line, " -> ")
		var lastX, lastY, xStart, yStart, xEnd, yEnd int
		for z, path := range paths {
			x, y, err := getCoords(path)

			if y > yBig {
				yBig = y
			}

			if err != nil {
				return nil, 0, err
			}

			if z == 0 {
				blocked[point.NewPoint2D(x, y)] = true
				lastX = x
				lastY = y
				continue
			}

			if x < lastX {
				xStart = x
				xEnd = lastX
			} else {
				xStart = lastX
				xEnd = x
			}
			if y < lastY {
				yStart = y
				yEnd = lastY
			} else {
				yStart = lastY
				yEnd = y
			}
			for j := yStart; j <= yEnd; j++ {
				for i := xStart; i <= xEnd; i++ {
					blocked[point.NewPoint2D(i, j)] = true
				}
			}
			lastX = x
			lastY = y
		}
	}

	return blocked, yBig, nil
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 14 --\n")
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
