package main

import (
	"bufio"
	"strconv"
	"strings"

	"../utils"
)

type node struct {
	start       int
	lenChildren int
	children    []node
	lenMetadata int
	metadata    []int
}

func main() {
	data := utils.GetProblemInput("08")
	defer data.Close()

	var numbers []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		for _, num := range s {
			dat, err := strconv.Atoi(num)
			utils.CheckError(err)
			numbers = append(numbers, dat)
		}
	}

	root := getNodeData(0, numbers)

	utils.PrintSolution(1, strconv.Itoa(addMetadata(root)))
	utils.PrintSolution(2, strconv.Itoa(addNodeValue(root)))
}

func getNodeData(start int, data []int) node {
	newNode := node{start: start,
		lenChildren: data[start],
		lenMetadata: data[start+1],
	}
	dataStart := start + 2

	for i := 0; i < newNode.lenChildren; i++ {
		child := getNodeData(dataStart, data)
		newNode.children = append(newNode.children, child)
		dataStart += getNodeLength(child)
	}

	for i := 0; i < newNode.lenMetadata; i++ {
		newNode.metadata = append(newNode.metadata, data[dataStart+i])
	}
	return newNode
}

func getNodeLength(n node) (result int) {
	result = 2 + n.lenMetadata
	for i := 0; i < n.lenChildren; i++ {
		result += getNodeLength(n.children[i])
	}
	return
}

func addMetadata(n node) (result int) {
	for _, child := range n.children {
		result += addMetadata(child)
	}
	for _, data := range n.metadata {
		result += data
	}
	return
}

func addNodeValue(root node) (result int) {
	if root.lenChildren == 0 {
		for _, dat := range root.metadata {
			result += dat
		}
	} else {
		for _, dat := range root.metadata {
			if dat > 0 && dat <= root.lenChildren {
				result += addNodeValue(root.children[dat-1])
			}
		}
	}
	return
}
