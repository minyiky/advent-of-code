package day03

import (
	"strconv"
	"strings"
)

func findJoltage(line string, codeLen int) (int, error) {
	code := ""
	for i := 0; i < codeLen; i++ {
		for _, v := range []string{"9", "8", "7", "6", "5", "4", "3", "2", "1"} {
			index := strings.Index(line[:len(line)-codeLen+i+1], v)
			if index == -1 {
				continue
			}
			line = line[index+1:]
			code += v
			break
		}
	}

	val, err := strconv.Atoi(code)
	if err != nil {
		return 0, err
	}
	return val, nil
}
