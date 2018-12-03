package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("../Inputs/Day02_Input.txt")
	check(err)
	defer data.Close()

	var s []string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	index1 := -1
	index2 := -1

	for i := 0; i < len(s)-1; i++ {
		diff := 0
		for j := i + 1; j < len(s); j++ {
			diff = 0
			for c := 0; c < len(s[i]); c++ {
				if s[i][c] != s[j][c] {
					diff++
				}
			}
			if diff == 1 {
				index2 = j
				break
			}
		}
		if diff == 1 {
			index1 = i
			break
		}
	}

	var buffer bytes.Buffer

	for i := 0; i < len(s[index1]); i++ {
		if s[index1][i] == s[index2][i] {
			buffer.WriteString(string(s[index1][i]))
		}
	}

	fmt.Print(buffer.String())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
