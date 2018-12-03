package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("../Inputs/Day02_Input.txt")
	check(err)
	defer data.Close()

	var s []string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	twos := 0
	threes := 0

	for _, code := range s {
		addTwo := true
		addThree := true
		var chars = make(map[string]int)
		for _, c := range code {
			if chars[string(c)] < 1 {
				chars[string(c)] = 1
			} else {
				chars[string(c)] += 1
			}
		}
		for _, value := range chars {
			if value == 2 && addTwo {
				twos++
				addTwo = false
			} else if value == 3 && addThree {
				threes++
				addThree = false
			}
		}
	}

	fmt.Print(twos * threes)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
