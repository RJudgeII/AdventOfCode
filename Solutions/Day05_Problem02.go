package main

import (
	"bufio"
	"fmt"
	"os"

	"./Packages/errors"
)

func main() {
	data, err := os.Open("../Inputs/Day05_Input.txt")
	errors.Check(err)
	defer data.Close()

	var s string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = scanner.Text()
	}

	r := []rune(s)

	minLen := len(r) + 1
	for cha := 65; cha <= 90; cha++ {
		newR := removeUnit(r, cha)
		for i := len(newR) - 2; i >= 0; i-- {
			if i != len(newR)-1 && checkPolarity(newR, i) {
				newR = append(newR[:i], newR[i+2:]...)
			}
		}
		if len(newR) < minLen {
			minLen = len(newR)
		}
	}

	fmt.Println(minLen)
}

func checkPolarity(r []rune, index int) (result bool) {
	ascii := int(r[index])
	ascii2 := int(r[index+1])
	result = false

	if ascii >= 65 && ascii <= 90 && ascii2 == ascii+32 {
		result = true
	} else if ascii >= 97 && ascii <= 122 && ascii2 == ascii-32 {
		result = true
	}

	return
}

func removeUnit(r []rune, val int) (result []rune) {
	result = make([]rune, len(r))
	copy(result, r)

	for i := len(result) - 1; i >= 0; i-- {
		if int(result[i]) == val || int(result[i]) == val+32 {
			result = append(result[:i], result[i+1:]...)
		}
	}

	return
}
