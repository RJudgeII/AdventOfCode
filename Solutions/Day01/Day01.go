package main

import (
	"bufio"
	"strconv"

	"../utils"
)

func main() {
	data := utils.GetProblemInput("01")
	defer data.Close()

	var s []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		utils.CheckError(err)
		s = append(s, i)
	}

	total := 0
	var freqs = make(map[int]bool)
	exists := false
	repeated := 0

	run := 1
	for {
		for _, diff := range s {
			total += diff
			if freqs[total] && !exists {
				exists = true
				repeated = total
			} else {
				freqs[total] = true
			}
		}
		if run == 1 {
			utils.PrintSolution(1, strconv.Itoa(total))
			run++
		}
		if exists {
			utils.PrintSolution(2, strconv.Itoa(repeated))
			break
		}
	}
}
