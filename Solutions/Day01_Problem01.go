package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("../Inputs/Day01_Input.txt")
	check(err)
	defer data.Close()

	var s []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		s = append(s, i)
	}

	out := 0

	for _, diff := range s {
		out += diff
	}

	fmt.Print(out)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
