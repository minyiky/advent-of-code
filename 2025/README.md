# 2025 Advent of Code solutions

![](https://img.shields.io/badge/tests%20passed%20🐹-23/24-important)
![](https://img.shields.io/badge/day%20📅-25-blue)
![](https://img.shields.io/badge/stars%20⭐-24-yellow)
![](https://img.shields.io/badge/days%20completed-12-red)

Here are my results for the [2025 advent of code](https://adventofcode.com/2021) competition


|              *Day*              | *Stars* |  *Solution*  |                         *Notes*                         |
|---------------------------------|---------|--------------|---------------------------------------------------------|
| Day 1: Secret Entrance |  ⭐⭐  | [go](day01/) | some interesting modular arithmetic today, a surprising number of places to get caught out by a day one puzzle |
| Day 2: Gift Shop |  ⭐⭐  | [go](day02/) | Pattern matching today, there is a nice regexp trick if your language support back referencing but unfortunately go does not. I took a slightly less brute force approach today by constructing the patterns and looking for matches rather than looping through all numbers |
| Day 3: Lobby |  ⭐⭐  | [go](day03/) | There were a few ways to attempt today's puzzle, I took the approach of looping from 9 down to 1 to find the first instance of the character and then substringing from there to create a smaller set. |
| Day 4: Printing Department |  ⭐⭐  | [go](day04/) | A classic grid problem today, part 2 was solvable by using part 1 in a loop with efficiencies made  by only check neighbours of removed stacks |
| Day 5: Cafeteria |  ⭐⭐  | [go](day05/) | A bit of additional processing at the start made part 2 a breeze today |
| Day 5: Trash Compacter |  ⭐⭐  | [go](day06/) | input parsing made today a bit more challenging but careful logic made it ok. A day that was actually easier in excel |
| Day 7: Laboratories |  ⭐⭐  | [go](day07/) | A classic BFS / DFS search today where memoisation was key. Benching revealved my DFS solution was ~25-50% quicker than my BFS solution |


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
