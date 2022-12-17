# 2022 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-32/32-success)
![](https://img.shields.io/badge/day%20📅-17-blue)
![](https://img.shields.io/badge/stars%20⭐-34-yellow)
![](https://img.shields.io/badge/days%20completed-17-red)

Here are my results for the [2022 advent of code](https://adventofcode.com/2022) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Calorie Counting         |  ⭐⭐  | [go](day01/) | Adding lists of numbers to find the highest             |
| Day 2: Rock, Paper, Scissors    |  ⭐⭐  | [go](day02/) | Using maps to compare strings for the game              |
| Day 3: Rucksack Reorganization  |  ⭐⭐  | [go](day03/) | Finding matching characters (runes in go)               |
| Day 4: Camp Cleanup             |  ⭐⭐  | [go](day04/) | Comparing ranges to find overlap                        |
| Day 5: Supply Stacks            |  ⭐⭐  | [go](day05/) | Parsing strings to find stacks then moving items around<br>This was more of a challenge today with string parsing and stack manipulation |
| Day 6: Tuning Trouble           |  ⭐⭐  | [go](day06/) | Sliding windows and contains checks                     |
| Day 7: No Space Left On Device  |  ⭐⭐  | [go](day07/) | Reading input commands for string manipulation          |
| Day 8: Treetop Tree House       |  ⭐⭐  | [go](day08/) | Searching through a 2d slice to find height values      |
| Day 9: Rope Bridge              |  ⭐⭐  | [go](day09/) | Vector comparisons to check for movement                |
| Day 10: Cathode-Ray Tube        |  ⭐⭐  | [go](day10/) | Tracking a variable set by text commands<br>The first day where part 2 was very different to part 1 |
| Day 11: Monkey in the Middle    |  ⭐⭐  | [go](day11/) | The hardest challenge so far, parsing complex data into a struct<br>Part 2 has a hidden challenge that you have to work out, big numbers can end up small... |
| Day 12: Hill Climbing Algorithm |  ⭐⭐  | [go](day12/) | Path finding through a set of heights<br>Used a map of distance for optimisation |
| Day 13: Distress Signal         |  ⭐⭐  | [go](day13/) | This was a day where the language choice made a big impact<br>A lot of effort was spent parsing the input while in python it was a single line `eval(line)`<br>The focus on parsing meant that I missed that I misunderstood the assignment RTFM! |
| Day 14: Regolith Reservoir      |  ⭐⭐  | [go](day14/) | A nice change today, nothing too bad, a classic "falling sand" problem |
| Day 15: Beacon Exclusion Zone   |  ⭐⭐  | [go](day15/) | Today was definitely a day where thinking was required<br>The simple map based approach I used at first was too slow in part 1<br>Checking all the positions was too slow in part 2 |
| Day 16: Proboscidea Volcanium   |  ⭐⭐  | [go](day16/) | Another tough day for the CPU, my initial solution was very unopimised and I stopped it after 10 minutes.<br>I eventually managed to get to <20s with a reduction of steps<br>code is messy and I need to investigate path finding algorithms |


## Running the code

To run the go code, you must be in this directory before running any of the following commands. You will also need to create `input.txt` files by running;
```
find ./day* -maxdepth 0 -type d -exec touch {}/example.txt \;
```

Any particular day can be executed by running:
```
go run ./dayXX/cmd/main.go
```

To run all days you can instead run:
```
go run ./cmd/runall/main.go
```

To run the tests use:
```
go test -v -count=1 ./...
```
