package day18

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

var directions = []aocutils.Vector3D{
	{X: 1, Y: 0, Z: 0},
	{X: -1, Y: 0, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: -1},
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
