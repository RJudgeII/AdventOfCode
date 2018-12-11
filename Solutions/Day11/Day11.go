package main

import (
	"bufio"
	"strconv"

	"../utils"
)

func main() {
	data := utils.GetProblemInput("11")
	defer data.Close()

	serial := 0
	var err error

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		serial, err = strconv.Atoi(scanner.Text())
		utils.CheckError(err)
	}

	grid := [][]int{}

	for x := 1; x <= 300; x++ {
		var col []int
		for y := 1; y <= 300; y++ {
			col = append(col, getPower(x, y, serial))
		}
		grid = append(grid, col)
	}

	maxPower, _ := getMaxPower(grid, 3)
	bigGrid := getLargestSquare(grid)

	utils.PrintSolution(1, maxPower)
	utils.PrintSolution(2, bigGrid)
}

func getPower(x, y, serial int) (result int) {
	result = x + 10
	result *= y
	result += serial
	result *= (x + 10)
	result = (result / 100) % 10
	result -= 5
	return
}

func getMaxPower(grid [][]int, size int) (result string, max int) {
	max = -100
	var x string
	var y string

	for xc := 0; xc < 301-size; xc++ {
		for yc := 0; yc < 301-size; yc++ {
			power := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					power += grid[xc+i][yc+j]
				}
			}
			if power > max {
				max = power
				x = strconv.Itoa(xc + 1)
				y = strconv.Itoa(yc + 1)
			}
		}
	}

	result = x + "," + y
	return
}

func getLargestSquare(grid [][]int) (result string) {
	power := -100

	for i := 1; i <= len(grid); i++ {
		coord, temp := getMaxPower(grid, i)
		if temp > power {
			power = temp
			result = coord + "," + strconv.Itoa(i)
		}
	}

	return
}
