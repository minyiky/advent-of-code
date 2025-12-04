# 2025 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-54/54-success)
![](https://img.shields.io/badge/day%20📅-25-blue)
![](https://img.shields.io/badge/stars%20⭐-39-yellow)
![](https://img.shields.io/badge/days%20completed-19-red)

Here are my results for the [2025 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Secret Entrance |  ⭐⭐  | [go](day01/) | some interesting modular arithmetic today, a surprising number of places to get caught out by a day one puzzle |
| Day 2: Gift Shop |  ⭐⭐  | [go](day02/) | Pattern matching today, there is a nice regexp trick if your language support back referencing but unfortunatley go does not. I took a slightly less brute force approach today by constructing the patterns and looking for matches rather than looping through all numbers |
| Day 3: Lobby |  ⭐⭐  | [go](day03/) | There were a few ways to attempt today's puzzle, I took the approach of looping from 9 down to 1 to find the first instance of the character and then sustringing from there to create a smaller set. |
| Day 4: Printing Department |  ⭐⭐  | [go](day04/) | A classic grid problem today, part 2 was solvable by using part 1 in a loop with efficiencies made  by only check neighbours of removed stacks |


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
