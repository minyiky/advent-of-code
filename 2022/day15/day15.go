package day15

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/maths"
	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

type Sensor struct {
	pos  point.Point2D
	dist int
}

func newSensor(x, y, dist int) *Sensor {
	return &Sensor{
		pos:  point.NewPoint2D(x, y),
		dist: dist,
	}
}

func (s *Sensor) InBounds(pos point.Point2D) bool {
	if pos == s.pos {
		return false
	}
	return point.ManhattanDistance(s.pos, pos) <= s.dist
}

func (s *Sensor) IgnoreNext(pos point.Point2D) int {
	yMax := s.dist - maths.Abs(pos.X()-s.pos.X())
	return yMax + (s.pos.Y() - pos.Y())
}

func mapBeacon(sensor, beacon point.Point2D, beacons, sensors map[point.Point2D]bool, topLeft, bottomRight *point.Point2D) *Sensor {
	dist := point.ManhattanDistance(sensor, beacon)
	if sensor.X()+dist > bottomRight.X() {
		bottomRight.SetX(sensor.X() + dist)
	}
	if sensor.X()-dist < topLeft.X() {
		topLeft.SetX(sensor.X() - dist)
	}
	if sensor.Y()+dist > topLeft.Y() {
		topLeft.SetY(sensor.Y() + dist)
	}
	if sensor.Y()-dist > bottomRight.Y() {
		bottomRight.SetY(sensor.Y() - dist)
	}
	beacons[beacon] = true
	sensors[sensor] = true
	return newSensor(sensor.X(), sensor.Y(), dist)
}

func mapBeaconNaive(sensor, beacon point.Point2D, coverage map[point.Point2D]bool, topLeft, bottomRight *point.Point2D) {
	dist := point.ManhattanDistance(sensor, beacon)
	if sensor.X()+dist > bottomRight.X() {
		bottomRight.SetX(sensor.X() + dist)
	}
	if sensor.X()-dist < topLeft.X() {
		topLeft.SetX(sensor.X() - dist)
	}
	if sensor.Y()+dist > topLeft.Y() {
		topLeft.SetY(sensor.Y() + dist)
	}
	if sensor.Y()-dist > bottomRight.Y() {
		bottomRight.SetY(sensor.Y() - dist)
	}
	for xLim := 0; xLim <= dist; xLim++ {
		y := dist - xLim
		for x := 0; x <= xLim; x++ {
			for _, pos := range []point.Point2D{
				point.NewPoint2D(sensor.X()+x, sensor.Y()+y),
				point.NewPoint2D(sensor.X()-x, sensor.Y()+y),
				point.NewPoint2D(sensor.X()+x, sensor.Y()-y),
				point.NewPoint2D(sensor.X()-x, sensor.Y()-y),
			} {
				if pos != beacon {
					coverage[pos] = true
				}
			}
		}
	}
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 15 --\n")
	if err := Part1(w, lines, 2000000); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	bottomLeft := point.NewPoint2D(0, 0)
	topRight := point.NewPoint2D(4000000, 4000000)
	if err := Part2(w, lines, &bottomLeft, &topRight); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
