package day13

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// TODO: Lots of tidying
func StripEnds(list string) string {
	return list[1 : len(list)-1]
}

func isValidSubsection(left, right []string) (bool, bool) {
	var end int
	if len(left) > len(right) {
		end = len(right)
	} else {
		end = len(left)
	}
	for i := 0; i < end; i++ {
		leftVal, _ := strconv.Atoi(left[i])
		rightVal, _ := strconv.Atoi(right[i])
		if rightVal < leftVal {
			return false, false
		}
		if rightVal > leftVal {
			return true, false
		}
	}
	if len(left) > len(right) {
		return false, false
	}
	if len(left) < len(right) {
		return true, false
	}
	return false, true
}

func Encapsulation(list string) string {
	var x int
	for i, char := range list {
		switch char {
		case '[':
			x++
		case ']':
			x--
			if x == 0 {
				return list[:i+1]
			}
		}
	}
	return ""
}

func Extend(list string, num int) (string, string) {
	index := strings.Index(list, ",")
	var val string
	if index == -1 {
		val = list
		list = ""
	} else {
		val = list[:index]
		list = list[index:]
	}
	newList := val
	for i := 1; i < num; i++ {
		newList += "," + val
	}
	return list, newList
}

func isValid(left, right string) (bool, bool) {
	var lil, ril bool // LeftIsList, RightIsList
	var substrLeft, substrRight string
	var leftList, rightList []string
	for left != "" && right != "" {
		lil = false
		ril = false

		if left[0] == ',' {
			left = left[1:]
		}
		if right[0] == ',' {
			right = right[1:]
		}
		if left[len(left)-1] == ',' {
			left = left[:len(left)-1]
		}
		if right[len(right)-1] == ',' {
			right = right[:len(right)-1]
		}

		if !strings.Contains(left, "[") && !strings.Contains(right, "[") {
			leftList := strings.Split(left, ",")
			rightList := strings.Split(right, ",")
			return isValidSubsection(leftList, rightList)
		}

		nextLeft := strings.Index(left, "[")
		nextRight := strings.Index(right, "[")

		if nextLeft == -1 {
			nextLeft = len(left)
		}
		if nextRight == -1 {
			nextRight = len(right)
		}

		if nextLeft == 0 {
			lil = true
			substrLeft = Encapsulation(left)
			left = left[len(substrLeft):]
			substrLeft = StripEnds(substrLeft)
		}

		if nextRight == 0 {
			ril = true
			substrRight = Encapsulation(right)
			right = right[len(substrRight):]
			substrRight = StripEnds(substrRight)
		}

		if lil || ril {
			if lil && !ril { // The left side is a list but the right isn't
				right, substrRight = Extend(right, 1)
				if ok, cont := isValid(substrLeft, substrRight); !cont {
					return ok, false
				}
				continue
			}
			if ril && !lil { // The right side is a list but the left isn't
				left, substrLeft = Extend(left, 1)
				if ok, cont := isValid(substrLeft, substrRight); !cont {
					return ok, false
				}
				continue
			}
			if ok, cont := isValid(substrLeft, substrRight); !cont {
				return ok, false
			}
			continue
		}

		var leftCount, rightCount []int

		for index, char := range left[:nextLeft] {
			if char == ',' {
				leftCount = append(leftCount, index)
			}
		}
		for index, char := range right[:nextRight] {
			if char == ',' {
				rightCount = append(rightCount, index)
			}
		}

		var rightIndex, leftIndex int
		if len(leftCount) == 0 || len(rightCount) == 0 {
			if len(leftCount) == 0 {
				leftIndex = nextLeft
			} else {
				leftIndex = leftCount[len(leftCount)-1]
			}
			if len(rightCount) == 0 {
				rightIndex = nextRight
			} else {
				rightIndex = rightCount[len(rightCount)-1]
			}
		} else if len(leftCount) > len(rightCount) {
			rightIndex = nextRight
			leftIndex = leftCount[len(rightCount)-1]
		} else if len(leftCount) < len(rightCount) {
			leftIndex = nextLeft
			rightIndex = rightCount[len(leftCount)-1]
		} else {
			rightIndex = rightCount[len(rightCount)-1]
			leftIndex = leftCount[len(leftCount)-1]
		}

		substrRight = right[:rightIndex]
		substrLeft = left[:leftIndex]
		right = right[rightIndex:]
		left = left[leftIndex:]

		if substrLeft[0] == ',' {
			substrLeft = substrLeft[1:]
		}
		if substrRight[0] == ',' {
			substrRight = substrRight[1:]
		}
		if substrLeft[len(substrLeft)-1] == ',' {
			substrLeft = substrLeft[:len(substrLeft)-1]
		}
		if substrRight[len(substrRight)-1] == ',' {
			substrRight = substrRight[:len(substrRight)-1]
		}

		leftList = strings.Split(substrLeft, ",")
		rightList = strings.Split(substrRight, ",")

		if ok, cont := isValidSubsection(leftList, rightList); !cont {
			return ok, false
		}
	}
	if left == "" && right == "" {
		return false, true
	}
	return left == "", false
}

func Run(w io.Writer) {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	fmt.Fprintf(w, "-- Solution for 2022 day 13 --\n")
	if err := Part1(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	if err := Part2(w, lines); err != nil {
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
