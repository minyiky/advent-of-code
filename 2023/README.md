# 2023 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-12/12-success)
![](https://img.shields.io/badge/day%20📅-6-blue)
![](https://img.shields.io/badge/stars%20⭐-12-yellow)
![](https://img.shields.io/badge/days%20completed-6-red)

Here are my results for the [2023 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Trebuchet?!              |  ⭐⭐  | [go](day01/) | This day was an exercise in string parsing to find digits written as digits and spelled out |
| Day 2: Cube Conundrum           |  ⭐⭐  | [go](day02/) | This was a day mainly about pattern matching with regex being a useful tool for this |
| Day 3: Gear Ratios              |  ⭐⭐  | [go](day03/) | One of the more tricky Day 3 challenges. Using a grid and map avoided extensive bounds checks |
| Day 4: Scratchcards             |  ⭐⭐  | [go](day04/) | Simple string matching exercise |
| Day 5: If You Give A Seed A Fertilizerards |  ⭐⭐  | [go](day05/) | A difficult day this early one with lots of room to turn your PC into a space heater, or run for a long time, some potential optimisations ended up being detrimental |
| Day 6: Wait for it              |  ⭐⭐  | [go](day06/) | It was time to bring out the GCSE maths today and solve a quadratic equation |
| Day 7: Camel Cards              |  ⭐⭐  | [go](day07/) | Poker with a twist, read the rules carefully and don't forget that you can have 5 of a kind |



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
