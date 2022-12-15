package day15

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

//go:embed input.txt
var input string

type Sensor struct {
	pos  aocutils.Vector
	dist int
}

func newSensor(x, y, dist int) *Sensor {
	return &Sensor{
		pos:  aocutils.NewVector(x, y),
		dist: dist,
	}
}

func (s *Sensor) InBounds(pos aocutils.Vector) bool {
	if pos == s.pos {
		return false
	}
	return s.pos.MDist(pos) <= s.dist
}

func (s *Sensor) IgnoreNext(pos aocutils.Vector) int {
	yMax := s.dist - aocutils.Abs(pos.X-s.pos.X)
	return yMax + (s.pos.Y - pos.Y)
}

func mapBeacon(sensor, beacon aocutils.Vector, beacons, sensors map[aocutils.Vector]bool, topLeft, bottomRight *aocutils.Vector) *Sensor {
	dist := sensor.MDist(beacon)
	if sensor.X+dist > bottomRight.X {
		bottomRight.X = sensor.X + dist
	}
	if sensor.X-dist < topLeft.X {
		topLeft.X = sensor.X - dist
	}
	if sensor.Y+dist > topLeft.Y {
		topLeft.Y = sensor.Y + dist
	}
	if sensor.Y-dist > bottomRight.Y {
		bottomRight.Y = sensor.Y - dist
	}
	beacons[beacon] = true
	sensors[sensor] = true
	return newSensor(sensor.X, sensor.Y, dist)
}

func mapBeaconNaive(sensor, beacon aocutils.Vector, coverage map[aocutils.Vector]bool, topLeft, bottomRight *aocutils.Vector) {
	dist := sensor.MDist(beacon)
	if sensor.X+dist > bottomRight.X {
		bottomRight.X = sensor.X + dist
	}
	if sensor.X-dist < topLeft.X {
		topLeft.X = sensor.X - dist
	}
	if sensor.Y+dist > topLeft.Y {
		topLeft.Y = sensor.Y + dist
	}
	if sensor.Y-dist > bottomRight.Y {
		bottomRight.Y = sensor.Y - dist
	}
	for xLim := 0; xLim <= dist; xLim++ {
		y := dist - xLim
		for x := 0; x <= xLim; x++ {
			for _, pos := range []aocutils.Vector{
				aocutils.NewVector(sensor.X+x, sensor.Y+y),
				aocutils.NewVector(sensor.X-x, sensor.Y+y),
				aocutils.NewVector(sensor.X+x, sensor.Y-y),
				aocutils.NewVector(sensor.X-x, sensor.Y-y),
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

	bottomLeft := aocutils.NewVector(0, 0)
	topRight := aocutils.NewVector(4000000, 4000000)
	if err := Part2(w, lines, bottomLeft, topRight); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
