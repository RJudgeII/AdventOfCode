package main

import (
	"bufio"
	"sort"
	"strconv"

	"../utils"
)

type cart struct {
	xPos      int
	yPos      int
	direction int
	turn      int
	id        int
}

func main() {
	data := utils.GetProblemInput("13")
	defer data.Close()

	carts := []cart{}
	tracks := []string{}

	lineNum := 0
	idNum := 0
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "^" {
				carts = append(carts, cart{i, lineNum, 0, 0, idNum})
				idNum++
				line = replace(line, '|', i)
			} else if string(line[i]) == ">" {
				carts = append(carts, cart{i, lineNum, 1, 0, idNum})
				idNum++
				line = replace(line, '-', i)
			} else if string(line[i]) == "v" {
				carts = append(carts, cart{i, lineNum, 2, 0, idNum})
				idNum++
				line = replace(line, '|', i)
			} else if string(line[i]) == "<" {
				carts = append(carts, cart{i, lineNum, 3, 0, idNum})
				idNum++
				line = replace(line, '-', i)
			}
		}
		tracks = append(tracks, line)
		lineNum++
	}

	collisions := 0
	xCol := -1
	yCol := -1
	oneLeft := false

	step := 0
	for {
		step++
		ind1 := -1
		ind2 := -1
		for k, c := range carts {
			if string(tracks[c.yPos][c.xPos]) == "+" {
				c.TurnAtIntersection()
			} else if string(tracks[c.yPos][c.xPos]) == "/" {
				c.TurnAtForwardCurve()
			} else if string(tracks[c.yPos][c.xPos]) == "\\" {
				c.TurnAtBackCurve()
			}
			c.Move()
			carts[k] = c
			if detectCollision(c, carts) {
				collisions++
				xCol = c.xPos
				yCol = c.yPos
				for j, car := range carts {
					if car.xPos == xCol && car.yPos == yCol {
						if ind1 == -1 {
							ind1 = j
						} else {
							ind2 = j
						}
					}
				}
				if collisions == 1 {
					utils.PrintSolution(1, strconv.Itoa(xCol)+","+strconv.Itoa(yCol))
				}
			}
		}
		if ind1 > -1 {
			carts = append(carts[:ind2], carts[ind2+1:]...)
			carts = append(carts[:ind1], carts[ind1+1:]...)
		}

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].yPos < carts[j].yPos {
				return true
			}
			if carts[i].yPos > carts[j].yPos {

			}
			return carts[i].xPos < carts[j].xPos
		})
		if len(carts) == 1 {
			oneLeft = true
		}
		if oneLeft {
			xCol = carts[0].xPos
			yCol = carts[0].yPos
			break
		}
	}

	utils.PrintSolution(2, strconv.Itoa(xCol)+","+strconv.Itoa(yCol))
}

func replace(s string, r rune, i int) string {
	out := []rune(s)
	out[i] = r
	return string(out)
}

func (car *cart) TurnAtIntersection() {
	if car.turn == 0 {
		car.direction--
		if car.direction == -1 {
			car.direction = 3
		}
	} else if car.turn == 2 {
		car.direction++
		if car.direction == 4 {
			car.direction = 0
		}
	}
	car.turn = (car.turn + 1) % 3
}

func (car *cart) TurnAtForwardCurve() {
	if car.direction == 0 || car.direction == 2 {
		car.direction++
	} else {
		car.direction--
	}
}

func (car *cart) TurnAtBackCurve() {
	if car.direction == 0 || car.direction == 2 {
		car.direction--
		if car.direction == -1 {
			car.direction = 3
		}
	} else {
		car.direction++
		if car.direction == 4 {
			car.direction = 0
		}
	}
}

func (car *cart) Move() {
	if car.direction == 0 {
		car.yPos--
	} else if car.direction == 1 {
		car.xPos++
	} else if car.direction == 2 {
		car.yPos++
	} else if car.direction == 3 {
		car.xPos--
	}
}

func detectCollision(car cart, carts []cart) (result bool) {
	inPos := 0
	x := car.xPos
	y := car.yPos
	for _, c := range carts {
		if c.xPos == x && c.yPos == y {
			inPos++
		}
	}
	if inPos > 1 {
		result = true
	} else {
		result = false
	}
	return
}
