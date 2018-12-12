package main

import (
	"bufio"
	"strconv"
	"strings"

	"../utils"
)

type marble struct {
	value    int
	next     *marble
	previous *marble
}

func main() {
	data := utils.GetProblemInput("09")
	defer data.Close()

	var s []string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")
	}

	numPlayers, err := strconv.Atoi(s[0])
	utils.CheckError(err)
	lastMarble, err := strconv.Atoi(s[6])
	utils.CheckError(err)

	utils.PrintSolution(1, strconv.Itoa(playGame(numPlayers, lastMarble)))
	utils.PrintSolution(1, strconv.Itoa(playGame(numPlayers, lastMarble*100)))
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

func getMax(players []int) (result int) {
	for _, val := range players {
		if val > result {
			result = val
		}
	}
	return
}

func playGame(numPlayers, lastMarble int) (result int) {
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

	result = getMax(players)
	return
}
