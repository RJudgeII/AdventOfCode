package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./Packages/errors"
)

func main() {
	data, err := os.Open("../Inputs/Day01_Input.txt")
	errors.Check(err)
	defer data.Close()

	var s []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		errors.Check(err)
		s = append(s, i)
	}

	out := 0

	for _, diff := range s {
		out += diff
	}

	fmt.Print(out)
}
