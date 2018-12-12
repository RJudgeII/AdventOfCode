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

	areas := make([]int, len(points))
	infinites := make(map[string]bool)
	infinites["."] = true

	for i := 0; i < gridMax.x; i++ {
		for j := 0; j < gridMax.y; j++ {
			dists[i][j] = gridMax.x + gridMax.y + 1
		}
	}

	for ind, p := range points {
		for i := 0; i < gridMax.x; i++ {
			for j := 0; j < gridMax.y; j++ {
				d := distance(p, Point{i, j})
				if d < dists[i][j] {
					dists[i][j] = d
					grid[i][j] = strconv.Itoa(ind)
				} else if d == dists[i][j] {
					grid[i][j] = "."
				}
			}
		}
	}

	for i := 0; i < gridMax.x; i++ {
		for j := 0; j < gridMax.y; j++ {
			closest := grid[i][j]
			if i == 0 || j == 0 || i == gridMax.x-1 || j == gridMax.y-1 {
				if !infinites[closest] {
					infinites[closest] = true
				}
			}
			if closest != "." {
				clInt, err := strconv.Atoi(closest)
				errors.Check(err)
				areas[clInt]++
			}
		}
	}

	maxArea := 0
	for i, a := range areas {
		if a > maxArea && !infinites[strconv.Itoa(i)] {
			maxArea = a
		}
	}

	fmt.Println(maxArea)
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
