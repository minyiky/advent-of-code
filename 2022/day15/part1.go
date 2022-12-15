package day15

import (
	"fmt"
	"io"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func Part1Val(lines []string, row int) (int, error) {
	var value int
	beaconMap := make(map[aocutils.Vector]bool)
	sensorMap := make(map[aocutils.Vector]bool)
	var topLeft, bottomRight aocutils.Vector
	sensors := make([]*Sensor, 0, len(lines))
	for _, line := range lines {
		var sensor, beacon aocutils.Vector
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		sensors = append(sensors, mapBeacon(sensor, beacon, beaconMap, sensorMap, &topLeft, &bottomRight))
	}

	for x := topLeft.X; x <= bottomRight.X; x++ {
		pos := aocutils.NewVector(x, row)
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
