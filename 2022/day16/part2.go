package day16

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/minyiky/advent-of-code-utils/pkg/container"
)

type Route struct {
	pressure int
	valves   []string
}

type RouteList []Route

func (r RouteList) Len() int           { return len(r) }
func (r RouteList) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RouteList) Less(i, j int) bool { return r[i].pressure < r[j].pressure }

func getBestRoute2(currentValve Valve, viableValves map[string]Valve, paths map[string]map[string]int, path []string, time, startPressure int, routeSlice RouteList) RouteList {
	if currentValve.name != "AA" {
		time += 1
	}

	if time < 26 {
		startPressure += currentValve.pressure * (26 - time)
		routeSlice = append(routeSlice, Route{
			pressure: startPressure,
			valves:   path,
		})
	}
	if time >= 26 || len(viableValves) == 0 {
		return routeSlice
	}

	for key, dest := range viableValves {
		mapCopy := container.CopyMap(viableValves)
		delete(mapCopy, key)
		pathCopy := container.CopySlice(path)
		pathCopy = append(pathCopy, key)
		pathLen := paths[currentValve.name][key]
		routeSlice = getBestRoute2(dest, mapCopy, paths, pathCopy, time+pathLen, startPressure, routeSlice)
	}
	return routeSlice
}

func Part2Val(lines []string) (int, error) {
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

	routes := make(RouteList, 0)
	routes = getBestRoute2(valves["AA"], valvesViable, paths, []string{}, 0, 0, routes)

	sort.Sort(routes)
	container.ReverseSlice(routes)
	highestP := 0

	for i, route := range routes[:routes.Len()-2] {
	elephant:
		for _, elephant := range routes[i+1:] {
			for _, valve := range route.valves {
				if _, ok := container.SliceContains(elephant.valves, valve); ok {
					continue elephant
				}
			}
			if route.pressure+elephant.pressure > highestP {
				highestP = route.pressure + elephant.pressure
			}
			break
		}
	}

	return highestP, nil
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
