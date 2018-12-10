package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./Packages/errors"
)

type node struct {
	start       int
	lenChildren int
	children    []node
	lenMetadata int
	metadata    []int
}

func main() {
	data, err := os.Open("../Inputs/Day08_Input.txt")
	errors.Check(err)
	defer data.Close()

	var numbers []int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		for _, num := range s {
			dat, err := strconv.Atoi(num)
			errors.Check(err)
			numbers = append(numbers, dat)
		}
	}

	root := getNodeData(0, numbers)
	fmt.Println(addMetadata(root))
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
