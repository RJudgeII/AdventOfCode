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

	for i := len(r) - 2; i >= 0; i-- {
		if i != len(r)-1 && checkPolarity(r, i) {
			r = append(r[:i], r[i+2:]...)
		}
	}

	fmt.Println(len(r))
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
