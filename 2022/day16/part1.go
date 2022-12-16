package day16

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code/2022/aocutils"
)

func getPath(current Valve, route map[string]int, valves map[string]Valve, dest string, len int) int {
	len += 1
	routes := sort.IntSlice{math.MaxInt}
	for _, tunnel := range current.tunnels {
		if tunnel == dest {
			return len
		}
		routeCopy := aocutils.CopyMap(route)
		if spot, ok := routeCopy[tunnel]; !ok || len < spot {
			routeCopy[tunnel] = len
			routes = append(routes, getPath(valves[tunnel], routeCopy, valves, dest, len))
		}
	}

	routes.Sort()
	return routes[0]
}

func getBestRoute(currentValve Valve, viableValves map[string]Valve, paths map[string]map[string]int, time, startPressure, maxPresure int) int {
	time += 1

	startPressure += currentValve.pressure * (30 - time)
	if time >= 30 || len(viableValves) == 0 {
		return startPressure
	}

	maxPressure := 0
	for key, dest := range viableValves {
		mapCopy := aocutils.CopyMap(viableValves)
		delete(mapCopy, key)
		pathLen := paths[currentValve.name][key]
		pressure := getBestRoute(dest, mapCopy, paths, time+pathLen, startPressure, maxPresure)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}

	return maxPressure
}

func Part1Val(lines []string) (int, error) {
	var value int
	valves := make(map[string]Valve)
	valvesViable := make(map[string]Valve)

	for _, line := range lines {
		var valveID, tunnelString string
		var flow int
		fmt.Sscanf(line, "Valve %s has flow rate=%d", &valveID, &flow, &tunnelString)
		tmp := strings.Split(line, "to valve")
		if tmp[1][0] == 's' {
			tmp[1] = tmp[1][1:]
		}
		tmp[1] = tmp[1][1:]
		tunnels := strings.Split(tmp[1], ", ")
		valves[valveID] = Valve{
			name:    valveID,
			tunnels: tunnels,
		}
		if flow > 0 {
			valvesViable[valveID] = Valve{
				name:     valveID,
				pressure: flow,
				tunnels:  tunnels,
			}
		}
	}

	paths := make(map[string]map[string]int)
	for start, valve := range valvesViable {
		tmpPaths := make(map[string]int)
		for end, _ := range valvesViable {
			if start == end {
				continue
			}
			emptyMap := map[string]int{
				start: 0,
			}
			tmpPaths[end] = getPath(valve, emptyMap, valves, end, 0)
		}
		paths[start] = tmpPaths
	}

	tmpPaths := make(map[string]int)
	for end, _ := range valvesViable {
		emptyMap := map[string]int{
			"AA": 0,
		}
		tmpPaths[end] = getPath(valves["AA"], emptyMap, valves, end, 0)

	}
	paths["AA"] = tmpPaths

	value = getBestRoute(valves["AA"], valvesViable, paths, -1, 0, 0)
	return value, nil
}

func Part1(w io.Writer, lines []string) error {
	start := time.Now()
	value, err := Part1Val(lines)
	if err != nil {
		return err
	}
	duration := time.Since(start)
	fmt.Fprintf(w, "You realise that at most you could release a total pressure of %d\n", value)
	fmt.Fprintf(w, "This took %.2fms\n", float64(duration)/1e6)
	return nil
}
