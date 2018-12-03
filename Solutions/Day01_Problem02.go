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

	var freqs = make(map[int]bool)

	freq := 0
	exists := false

	for {
		for _, diff := range s {
			freq += diff
			if freqs[freq] {
				exists = true
				break
			} else {
				freqs[freq] = true
			}
		}
		if exists {
			break
		}
	}

	fmt.Print(freq)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
