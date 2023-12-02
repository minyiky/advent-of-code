# 2023 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-4/4-success)
![](https://img.shields.io/badge/day%20📅-2-blue)
![](https://img.shields.io/badge/stars%20⭐-4-yellow)
![](https://img.shields.io/badge/days%20completed-2-red)

Here are my results for the [2023 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Trebuchet?!              |  ⭐⭐  | [go](day01/) | This day was an exercise in string parsing to find digits written as digits and spelled out |



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
