package main

import (
	"bufio"
	"strconv"
	"strings"

	"../utils"
)

func main() {
	data := utils.GetProblemInput("03")
	defer data.Close()

	var left []int
	var top []int
	var width []int
	var height []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		le, err := strconv.Atoi(strings.Split(s[2], ",")[0])
		utils.CheckError(err)
		left = append(left, le)

		to, err := strconv.Atoi(strings.Split(strings.Split(s[2], ",")[1], ":")[0])
		utils.CheckError(err)
		top = append(top, to)

		wi, err := strconv.Atoi(strings.Split(s[3], "x")[0])
		utils.CheckError(err)
		width = append(width, wi)

		he, err := strconv.Atoi(strings.Split(s[3], "x")[1])
		utils.CheckError(err)
		height = append(height, he)
	}

	cloth := [1000][1000]int{}

	for i := 0; i < len(left); i++ {
		for l := left[i]; l < left[i]+width[i]; l++ {
			for t := top[i]; t < top[i]+height[i]; t++ {
				cloth[l][t]++
			}
		}
	}

	count := 0

	for _, row := range cloth {
		for _, col := range row {
			if col > 1 {
				count++
			}
		}
	}

	id := 0

	for i := 0; i < len(left); i++ {
		overlap := false
		for l := left[i]; l < left[i]+width[i]; l++ {
			for t := top[i]; t < top[i]+height[i]; t++ {
				if cloth[l][t] > 1 {
					overlap = true
				}
			}
		}
		if !overlap {
			id = i + 1
			break
		}
	}

	utils.PrintSolution(1, strconv.Itoa(count))
	utils.PrintSolution(2, strconv.Itoa(id))
}
