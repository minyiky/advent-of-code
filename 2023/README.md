# 2023 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-40/40-success)
![](https://img.shields.io/badge/day%20📅-20-blue)
![](https://img.shields.io/badge/stars%20⭐-38-yellow)
![](https://img.shields.io/badge/days%20completed-19-red)

Here are my results for the [2023 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Trebuchet?!               |  ⭐⭐  | [go](day01/) | This day was an exercise in string parsing to find digits written as digits and spelled out |
| Day 2: Cube Conundrum            |  ⭐⭐  | [go](day02/) | This was a day mainly about pattern matching with regex being a useful tool for this |
| Day 3: Gear Ratios               |  ⭐⭐  | [go](day03/) | One of the more tricky Day 3 challenges. Using a grid and map avoided extensive bounds checks |
| Day 4: Scratchcards              |  ⭐⭐  | [go](day04/) | Simple string matching exercise |
| Day 5: If You Give A Seed A Fertilizer |  ⭐⭐  | [go](day05/) | A difficult day this early one with lots of room to turn your PC into a space heater, or run for a long time, some potential optimisations ended up being detrimental |
| Day 6: Wait for it               |  ⭐⭐  | [go](day06/) | It was time to bring out the GCSE maths today and solve a quadratic equation |
| Day 7: Camel Cards               |  ⭐⭐  | [go](day07/) | Poker with a twist, read the rules carefully and don't forget that you can have 5 of a kind |
| Day 8: Haunted Wasteland         |  ⭐⭐  | [go](day08/) | Day 8 was generally a simple route following exercise but part two required inspection of input for a hidden constraint |
| Day 9: Mirage Maintenance        |  ⭐⭐  | [go](day09/) | a simple recursive formula with little change required for part 2 |
| Day 10: Pipe Maze                |  ⭐⭐  | [go](day10/) | Following the pipes was easy but determining the contents had a few tricks for switching the location, a good day for visualisations |
| Day 11: Cosmic Expansion         |  ⭐⭐  | [go](day11/) | By modifying grid values rather than duplicating values the day solves quickly |
| Day 12: Hot Springs              |  ⭐⭐  | [go](day12/) | As is common part one was brute forcable by combinatorics but that would not work for part two, instead a [dp](https://stackoverflow.blog/2022/01/31/the-complete-beginners-guide-to-dynamic-programming/) appraoch needed to be taken |
| Day 13: Point of Incidence       |  ⭐⭐  | [go](day13/) | Another day where parsing the instructions was more difficult than the puzzle, a lot of people got caught out in part 2 |
| Day 14: Parabolic Reflector Dish |  ⭐⭐  | [go](day14/) | Rotations and directions caught me out for a bit, and brute forcing again won't work, take time to print out a subsection of results to spot the trick |
| Day 15: Lens Library             |  ⭐⭐  | [go](day15/) | One of the easier days to code but the part instructions took a while to parse. More use of golang runes for calculations |
| Day 16: The Floor Will Be Lava   |  ⭐⭐  | [go](day16/) | Another "pipe" day with routes to follow, part 2 was a natural extension of part 1. Watch out for infinites |
| Day 17: Clumsy Crucible          |  ⭐⭐  | [go](day17/) | Today was the time to break out the better search algorithms. Todays solution was heavily inspired by [mnml](https://github.com/mnml/aoc/blob/main/2023/17/1.go) |


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
