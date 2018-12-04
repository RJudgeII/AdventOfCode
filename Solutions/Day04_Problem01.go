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

type TimeTotals struct {
	minutes [60]int
	total   int
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
	var guardTimes = make(map[string]*TimeTotals)

	var start int
	var end int

	for _, point := range guards {
		switcher := point.datum
		if string(switcher[0]) == "#" {
			if _, ok := guardTimes[switcher]; !ok {
				guardTimes[switcher] = &TimeTotals{[60]int{}, 0}
			}
			guardID = switcher
		} else if switcher == "asleep" {
			start = point.dateTime.Minute()
		} else if switcher == "up" {
			end = point.dateTime.Minute()
			for i := start; i < end; i++ {
				guardTimes[guardID].minutes[i] += 1
			}
			guardTimes[guardID].total += (end - start)
		}
	}

	maxTotal := -1
	for id, val := range guardTimes {
		if val.total > maxTotal {
			maxTotal = val.total
			guardID = id
		}
	}

	minute := -1
	maxTotal = -1
	for id, val := range guardTimes[guardID].minutes {
		if val > maxTotal {
			minute = id
			maxTotal = val
		}
	}

	idNum, err := strconv.Atoi(strings.Split(guardID, "#")[1])
	errors.Check(err)

	fmt.Println(idNum * minute)
}
