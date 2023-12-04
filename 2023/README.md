# 2023 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-6/6-success)
![](https://img.shields.io/badge/day%20📅-3-blue)
![](https://img.shields.io/badge/stars%20⭐-6-yellow)
![](https://img.shields.io/badge/days%20completed-3-red)

Here are my results for the [2023 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Trebuchet?!              |  ⭐⭐  | [go](day01/) | This day was an exercise in string parsing to find digits written as digits and spelled out |
| Day 2: Cube Conundrum           |  ⭐⭐  | [go](day02/) | This was a day mainly about pattern matching with regex being a useful tool for this |
| Day 3: Gear Ratios              |  ⭐⭐  | [go](day03/) | One of the more tricky Day 3 challenges. Using a grid and map avoided extensive bounds checks |
| Day 4: Scratchcards             |  ⭐⭐  | [go](day04/) | Simple string matching exercise |



## Running the code

To run the go code, you must be in this directory before running any of the following commands. You will also need to create `input.txt` files. To create an input for the current day run:

``` bash
make input
```

The current day can be executed by running:
``` bash
make run
```
or any specific day can run using
```bash
make run DAY=<DD>
```

To run all days you can instead run:
```bash
make run-all
```

To run the tests use:
```bash
make test
```
