# 2021 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-30/30-success)
![](https://img.shields.io/badge/day%20📅-25-blue)
![](https://img.shields.io/badge/stars%20⭐-50-yellow)
![](https://img.shields.io/badge/days%20completed-25-red)

Here are my results for the [2021 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1:  |  ⭐⭐  | [go](day01/) | |
| Day 2:  |  ⭐⭐  | [go](day02/) | |
| Day 3:  |  ⭐⭐  | [go](day03/) | |
| Day 4:  |  ⭐⭐  | [go](day04/) | |
| Day 5:  |  ⭐⭐  | [go](day05/) | |
| Day 6:  |  ⭐⭐  | [go](day06/) | |
| Day 7:  |  ⭐⭐  | [go](day07/) | |
| Day 8:  |  ⭐⭐  | [go](day08/) | |
| Day 9:  |  ⭐⭐  | [go](day09/) | |
| Day 10: |  ⭐⭐  | [go](day10/) | |
| Day 11: |  ⭐⭐  | [go](day11/) | |
| Day 12: |  ⭐⭐  | [go](day12/) | |
| Day 13: |  ⭐⭐  | [go](day13/) | |
| Day 14: |  ⭐⭐  | [go](day14/) | |
| Day 15: |  ⭐⭐  | [go](day15/) | |



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
