package main

import (
	"bufio"
	"strconv"
	"strings"

	"../utils"
)

func main() {
	data := utils.GetProblemInput("12")
	defer data.Close()

	padding := ""
	for i := 0; i < 22; i++ {
		padding += "."
	}
	potList := padding
	var condition = make(map[string]string)

	lineNum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if lineNum == 0 {
			potList += strings.Split(scanner.Text(), " ")[2]
			potList += padding
		} else if lineNum > 1 {
			condition[strings.Split(scanner.Text(), " => ")[0]] = strings.Split(scanner.Text(), " => ")[1]
		}
		lineNum++
	}

	for step := 0; step < 20; step++ {
		potList = getNextGeneration(potList, condition)
	}

	utils.PrintSolution(1, strconv.Itoa(getScore(potList, padding)))

	step := 0
	lastScore := 0
	var diffs []int
	for {
		step++
		potList = getNextGeneration(potList, condition)
		score := getScore(potList, padding)
		diffs = append(diffs, score-lastScore)
		lastScore = score

		if len(diffs) > 3 {
			if diffs[len(diffs)-1] == diffs[len(diffs)-2] && diffs[len(diffs)-2] == diffs[len(diffs)-3] {
				break
			}
		}
	}

	finalScore := lastScore + (50000000000-step-20)*diffs[len(diffs)-1]

	utils.PrintSolution(2, strconv.Itoa(finalScore))
}

func getNextGeneration(potList string, condition map[string]string) (result string) {
	result = ".."
	for i := 2; i < len(potList)-2; i++ {
		sub := potList[i-2 : i+3]
		result += condition[sub]
	}
	result += "..."
	return
}

func getScore(potList, padding string) (result int) {
	result = 0
	for i := 0; i < len(potList); i++ {
		if string(potList[i]) == "#" {
			result += i - len(padding)
		}
	}
	return
}
