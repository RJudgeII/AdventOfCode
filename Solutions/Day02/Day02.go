package main

import (
	"bufio"
	"bytes"
	"strconv"

	"../utils"
)

func main() {
	data := utils.GetProblemInput("02")
	defer data.Close()

	var s []string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	twos := 0
	threes := 0
	index1 := -1
	index2 := -2

	for i := 0; i < len(s); i++ {
		addTwo := true
		addThree := true
		chars := make(map[string]int)
		for _, c := range s[i] {
			if chars[string(c)] < 1 {
				chars[string(c)] = 1
			} else {
				chars[string(c)]++
			}
		}
		for _, val := range chars {
			if val == 2 && addTwo {
				twos++
				addTwo = false
			} else if val == 3 && addThree {
				threes++
				addThree = false
			}
		}
		if i < len(s)-1 {
			diff := 0
			for j := i + 1; j < len(s); j++ {
				diff = 0
				for c := 0; c < len(s[i]); c++ {
					if s[i][c] != s[j][c] {
						diff++
					}
				}
				if diff == 1 {
					index2 = j
					break
				}
			}
			if diff == 1 {
				index1 = i
			}
		}
	}

	var buffer bytes.Buffer

	for i := 0; i < len(s[index1]); i++ {
		if s[index1][i] == s[index2][i] {
			buffer.WriteString(string(s[index1][i]))
		}
	}

	utils.PrintSolution(1, strconv.Itoa(twos*threes))
	utils.PrintSolution(2, buffer.String())
}
