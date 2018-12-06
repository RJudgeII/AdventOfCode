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

type Point struct {
	x int
	y int
}

func main() {
	data, err := os.Open("../Inputs/Day06_Input.txt")
	errors.Check(err)
	defer data.Close()

	var points []Point

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ", ")

		xVal, err := strconv.Atoi(s[0])
		errors.Check(err)

		yVal, err := strconv.Atoi(s[1])
		errors.Check(err)

		points = append(points, Point{xVal, yVal})
	}

	gridMax := maxGridSize(points)

	grid := make([][]string, gridMax.x)
	for i := range grid {
		grid[i] = make([]string, gridMax.y)
	}
	dists := make([][]int, gridMax.x)
	for i := range dists {
		dists[i] = make([]int, gridMax.y)
	}

	for i := 0; i < gridMax.x; i++ {
		for j := 0; j < gridMax.y; j++ {
			dists[i][j] = allDists(Point{i, j}, points)
		}
	}

	area := 0
	for i := 0; i < gridMax.x; i++ {
		for j := 0; j < gridMax.y; j++ {
			if dists[i][j] < 10000 {
				area++
			}
		}
	}

	fmt.Println(area)
}

func maxGridSize(p []Point) (result Point) {
	x := -1
	y := -1
	for _, val := range p {
		if val.x > x {
			x = val.x
		}
		if val.y > y {
			y = val.y
		}
	}
	result = Point{x, y}
	return
}

func distance(p, q Point) (result int) {
	result = int(math.Abs(float64(q.x-p.x)) + math.Abs(float64(q.y-p.y)))
	return
}

func allDists(q Point, ps []Point) (result int) {
	result = 0
	for _, p := range ps {
		result += distance(q, p)
	}
	return
}
