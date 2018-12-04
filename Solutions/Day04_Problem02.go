package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"./Packages/errors"
)

type GuardData struct {
	dateTime time.Time
	datum    string
}

type Minutes struct {
	minutes [60]int
}

func main() {
	data, err := os.Open("../Inputs/Day04_Input.txt")
	errors.Check(err)
	defer data.Close()

	var guards []GuardData

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		givenDate := strings.Split(s[0], "[")[1] + " " + strings.Split(s[1], "]")[0]
		dat, err := time.Parse("2006-01-02 15:04", givenDate)
		errors.Check(err)

		swi := s[3]

		guards = append(guards, GuardData{dat, swi})
	}

	sort.Slice(guards, func(i, j int) bool { return guards[j].dateTime.After(guards[i].dateTime) })

	var guardID string
	var guardTimes = make(map[string]*Minutes)

	var start int
	var end int

	for _, point := range guards {
		switcher := point.datum
		if string(switcher[0]) == "#" {
			if _, ok := guardTimes[switcher]; !ok {
				guardTimes[switcher] = &Minutes{[60]int{}}
			}
			guardID = switcher
		} else if switcher == "asleep" {
			start = point.dateTime.Minute()
		} else if switcher == "up" {
			end = point.dateTime.Minute()
			for i := start; i < end; i++ {
				guardTimes[guardID].minutes[i] += 1
			}
		}
	}

	maxMinute := -1
	minute := -1

	for id, arr := range guardTimes {
		//fmt.Println(arr)
		for ind, min := range arr.minutes {
			if min > maxMinute {
				guardID = id
				maxMinute = min
				minute = ind
			}
		}
	}

	idNum, err := strconv.Atoi(strings.Split(guardID, "#")[1])
	errors.Check(err)

	fmt.Println(idNum * minute)
}
