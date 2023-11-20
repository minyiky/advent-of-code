package day18

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/minyiky/advent-of-code-utils/pkg/point"
)

//go:embed input.txt
var input string

var directions = []point.Point3D{
	point.NewPoint3D(1, 0, 0),
	point.NewPoint3D(-1, 0, 0),
	point.NewPoint3D(0, 1, 0),
	point.NewPoint3D(0, -1, 0),
	point.NewPoint3D(0, 0, 1),
	point.NewPoint3D(0, 0, -1),
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 18 --\n")
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
