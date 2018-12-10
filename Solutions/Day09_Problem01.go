package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./Packages/errors"
)

func main() {
	data, err := os.Open("../Inputs/Day09_Input.txt")
	errors.Check(err)
	defer data.Close()

	var s []string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")
	}

	numPlayers, err := strconv.Atoi(s[0])
	errors.Check(err)
	lastMarble, err := strconv.Atoi(s[6])
	errors.Check(err)

	players := make([]int, numPlayers)

	var marbles []int
	marbles = append(marbles, 0)

	index := 1
	for i := 1; i <= lastMarble; i++ {
		if i%23 != 0 {
			index += 2
			if index > len(marbles) {
				index = 1
			}
			marbles = placeMarble(i, index, marbles)
		} else {
			player := (i - 1) % numPlayers
			players[player] += i
			index -= 7
			if index < 0 {
				index += len(marbles)
			}
			var value int
			value, marbles = removeMarble(index, marbles)
			players[player] += value
		}
	}

	max := 0
	for _, val := range players {
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

func placeMarble(value int, ind int, marbles []int) (outMarbles []int) {
	if ind == len(marbles) {
		outMarbles = append(marbles, value)
	} else {
		outMarbles = append(marbles, 0)
		copy(outMarbles[ind+1:], outMarbles[ind:])
		outMarbles[ind] = value
	}
	return
}

func removeMarble(index int, marbles []int) (value int, outMarbles []int) {
	value = marbles[index]
	outMarbles = append(marbles[:index], marbles[index+1:]...)
	return
}
