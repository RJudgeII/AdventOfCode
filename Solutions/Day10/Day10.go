package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"../utils"
)

type star struct {
	x  int
	y  int
	vX int
	vY int
}

func main() {
	data := utils.GetProblemInput("10")
	defer data.Close()

	var stars []star

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		x, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[10:16]))
		utils.CheckError(err)

		y, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[18:24]))
		utils.CheckError(err)

		vX, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[36:38]))
		utils.CheckError(err)

		vY, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[40:42]))
		utils.CheckError(err)

		stars = append(stars, star{x, y, vX, vY})
	}

	step := 0
	area := 0
	lastArea := math.MaxInt64

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	lastMinX := 0
	lastMaxX := 0
	lastMinY := 0
	lastMaxY := 0

	for {
		step++
		minX = math.MaxInt32
		maxX = -1 * math.MaxInt32
		minY = math.MaxInt32
		maxY = -1 * math.MaxInt32
		for _, sta := range stars {
			x := sta.x + sta.vX*step
			y := sta.y + sta.vY*step

			if x > maxX {
				maxX = x
			} else if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			} else if y < minY {
				minY = y
			}
		}
		area = (maxX - minX) * (maxY - minY)
		if area > lastArea {
			break
		}
		lastArea = area
		lastMinX = minX
		lastMaxX = maxX
		lastMinY = minY
		lastMaxY = maxY
	}
	step--

	starMap := make([][]bool, lastMaxY-lastMinY+1)
	for i := 0; i < len(starMap); i++ {
		starMap[i] = make([]bool, lastMaxX-lastMinX+1)
	}

	for _, sta := range stars {
		starMap[sta.y+sta.vY*step-lastMinY][sta.x+sta.vX*step-lastMinX] = true
	}

	output := ""
	for i := 0; i < len(starMap); i++ {
		for j := 0; j < len(starMap[0]); j++ {
			if starMap[i][j] {
				output += "#"
			} else {
				output += " "
			}
		}
		output += "\n"
	}

	utils.PrintSolution(1, output)
	utils.PrintSolution(2, strconv.Itoa(step))
}
