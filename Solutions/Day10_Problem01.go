package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"./Packages/errors"
)

type star struct {
	x  int
	y  int
	vX int
	vY int
}

func main() {
	data, err := os.Open("../Inputs/Day10_Input.txt")
	errors.Check(err)
	defer data.Close()

	var stars []star

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		x, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[10:16]))
		errors.Check(err)

		y, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[18:24]))
		errors.Check(err)

		vX, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[36:38]))
		errors.Check(err)

		vY, err := strconv.Atoi(strings.TrimSpace(scanner.Text()[40:42]))
		errors.Check(err)

		stars = append(stars, star{x, y, vX, vY})
	}

	step := 0
	area := 0
	lastArea := math.MaxInt64

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

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
	}
	step--

	starMap := make([][]bool, maxY-minY+1)
	for i := 0; i < len(starMap); i++ {
		starMap[i] = make([]bool, maxX-minX+1)
	}

	for _, sta := range stars {
		starMap[sta.y+sta.vY*step-minY][sta.x+sta.vX*step-minX] = true
	}

	for i := 0; i < len(starMap); i++ {
		for j := 0; j < len(starMap[0]); j++ {
			if starMap[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
