package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/minyiky/advent-of-code/2022/day01"
	"github.com/minyiky/advent-of-code/2022/day02"
	"github.com/minyiky/advent-of-code/2022/day03"
	"github.com/minyiky/advent-of-code/2022/day04"
	"github.com/minyiky/advent-of-code/2022/day05"
	"github.com/minyiky/advent-of-code/2022/day06"
	"github.com/minyiky/advent-of-code/2022/day07"
	"github.com/minyiky/advent-of-code/2022/day08"
	"github.com/minyiky/advent-of-code/2022/day09"
	"github.com/minyiky/advent-of-code/2022/day10"
)

var DayMap = map[string]func(io.Writer){
	"day01": day01.Run,
	"day02": day02.Run,
	"day03": day03.Run,
	"day04": day04.Run,
	"day05": day05.Run,
	"day06": day06.Run,
	"day07": day07.Run,
	"day08": day08.Run,
	"day09": day09.Run,
	"day10": day10.Run,
}

func main() {
	http.HandleFunc("/", handleMain)

	for i := 0; i <= 25; i++ {
		index := fmt.Sprintf("%02d", i)
		http.HandleFunc("/day"+index, handleDay)
	}

	http.ListenAndServe(":8080", nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the main page!\n")
}

func handleDay(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Path[1:]

	dayFunc, ok := DayMap[day]
	if !ok {
		fmt.Fprintf(w, "I haven't implimented a solution for this day yet!\n")
		return
	}

	dayFunc(w)
}
