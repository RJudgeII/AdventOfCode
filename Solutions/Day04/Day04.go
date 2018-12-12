package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type guardData struct {
	dateTime time.Time
	datum    string
}

type timeTotals struct {
	minutes [60]int
	total   int
}

func main() {
	data := utils.GetProblemInput("04")
	defer data.Close()

	var guards []guardData

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		givenDate := strings.Split(s[0], "[")[1] + " " + strings.Split(s[1], "]")[0]
		dat, err := time.Parse("2006-01-02 15:04", givenDate)
		utils.CheckError(err)

		swi := s[3]

		guards = append(guards, guardData{dat, swi})
	}

	sort.Slice(guards, func(i, j int) bool { return guards[j].dateTime.After(guards[i].dateTime) })

	var guardID string
	var guardTimes = make(map[string]*timeTotals)

	var start int
	var end int

	for _, point := range guards {
		switcher := point.datum
		if string(switcher[0]) == "#" {
			if _, ok := guardTimes[switcher]; !ok {
				guardTimes[switcher] = &timeTotals{[60]int{}, 0}
			}
			guardID = switcher
		} else if switcher == "asleep" {
			start = point.dateTime.Minute()
		} else if switcher == "up" {
			end = point.dateTime.Minute()
			for i := start; i < end; i++ {
				guardTimes[guardID].minutes[i]++
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
	utils.CheckError(err)

	utils.PrintSolution(1, strconv.Itoa(idNum*minute))

	maxMinute := -1
	minute = -1

	for id, arr := range guardTimes {
		for ind, min := range arr.minutes {
			if min > maxMinute {
				guardID = id
				maxMinute = min
				minute = ind
			}
		}
	}

	idNum, err = strconv.Atoi(strings.Split(guardID, "#")[1])
	utils.CheckError(err)

	utils.PrintSolution(2, strconv.Itoa(idNum*minute))

}
