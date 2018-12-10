package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./Packages/errors"
)

type marble struct {
	value    int
	next     *marble
	previous *marble
}

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
	lastMarble *= 100

	players := make([]int, numPlayers)

	current := &marble{value: 0}
	current.next = &marble{1, current, current}
	current.previous = current.next

	for nextMarble := 2; nextMarble < lastMarble; nextMarble++ {
		if nextMarble%23 != 0 {
			newMarble := &marble{nextMarble, current.next.next, current.next}
			newMarble.previous.next = newMarble
			newMarble.next.previous = newMarble
			current = newMarble
		} else {
			player := (nextMarble - 1) % len(players)
			players[player] += nextMarble
			for i := 0; i < 7; i++ {
				current = current.previous
			}
			removed := current
			removed.next.previous = removed.previous
			removed.previous.next = removed.next
			players[player] += removed.value
			current = removed.next
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
