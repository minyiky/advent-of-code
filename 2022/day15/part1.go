package day15

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

func Part1Val(lines []string, row int) (int, error) {
	var value int
	beaconMap := make(map[point.Point2D]bool)
	sensorMap := make(map[point.Point2D]bool)
	var topLeft, bottomRight point.Point2D
	sensors := make([]*Sensor, 0, len(lines))
	for _, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		sensor := point.NewPoint2D(sensorX, sensorY)
		beacon := point.NewPoint2D(beaconX, beaconY)
		sensors = append(sensors, mapBeacon(sensor, beacon, beaconMap, sensorMap, &topLeft, &bottomRight))
	}

	for x := topLeft.X(); x <= bottomRight.X(); x++ {
		pos := point.NewPoint2D(x, row)
		if !beaconMap[pos] && !sensorMap[pos] {
			for _, s := range sensors {
				if s.InBounds(pos) {
					value++
					break
				}
			}
		}
	}
	return value, nil
}

func Part1(w io.Writer, lines []string, row int) error {
	start := time.Now()
	value, err := Part1Val(lines, row)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "When looking at positions with a y coordinate of %d, %d posisitions were covered by the sensors\n", row, value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
