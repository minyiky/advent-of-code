package day15

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part2Val(lines []string, bottomLeft, topRight *point.Point2D) (int, error) {
	var value int
	beaconMap := make(map[point.Point2D]bool)
	sensorMap := make(map[point.Point2D]bool)
	var topLeft, bottomRight point.Point2D
	sensors := make([]*Sensor, 0, len(lines))
	for _, line := range lines {
		var sensor, beacon point.Point2D
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", sensor.X(), sensor.Y(), beacon.X(), beacon.Y())
		sensors = append(sensors, mapBeacon(sensor, beacon, beaconMap, sensorMap, &topLeft, &bottomRight))
	}

	for x := bottomLeft.X(); x <= topRight.X(); x++ {
	loop:
		for y := bottomLeft.Y(); y <= topRight.Y(); y++ {
			pos := point.NewPoint2D(x, y)
			if !beaconMap[pos] && !sensorMap[pos] {
				for _, s := range sensors {
					if s.InBounds(pos) {
						y += s.IgnoreNext(pos)
						continue loop
					}
				}
				value = x*4000000 + y
			}
		}
	}
	return value, nil
}

func Part2(w io.Writer, lines []string, bottomLeft, topRight *point.Point2D) error {
	start := time.Now()
	value, err := Part2Val(lines, bottomLeft, topRight)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "With a more resricted are to look at the distress beacon must have a tuning frequency of %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
