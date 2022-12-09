package day07

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func ChangeDir(dir, cmd string) string {
	if cmd == "/" {
		return "/"
	}

	for _, subDir := range strings.Split(cmd, "/") {
		if subDir == ".." {
			index := strings.LastIndex(dir, "/")
			dir = dir[:index]
			continue
		}
		if dir == "/" {
			dir += subDir
			continue
		}

		dir += "/" + subDir
	}

	return dir
}

func ReadDirs(lines []string) map[string]int {
	dirs := make(map[string]int)
	currentDir := ""

	for _, line := range lines {
		if line[0:1] == "$" {
			cmds := strings.Split(line, " ")[1:]
			if cmds[0] == "cd" {
				currentDir = ChangeDir(currentDir, cmds[1])
				continue
			}
		}

		parts := strings.Split(line, " ")

		if size, err := strconv.Atoi(parts[0]); err == nil {
			dirs["/"] += size
			if len(currentDir) == 1 {
				continue
			}
			for i := 1; i < len(strings.Split(currentDir, "/")); i++ {
				dir := strings.Join(strings.Split(currentDir, "/")[:i+1], "/")
				dirs[dir] += size
			}
			continue
		}
	}

	return dirs
}

func Run() {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	log.Println("-- Solution for 2022 day 07 --")
	Part1(lines)
	Part2(lines)
}
