package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"./Packages/errors"
)

func main() {
	data, err := os.Open("../Inputs/Day07_Input.txt")
	errors.Check(err)
	defer data.Close()

	precedents := make(map[string][]string)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		fir := s[1]
		las := s[7]

		if len(precedents[las]) == 0 {
			precedents[las] = []string{fir}
		} else {
			precedents[las] = append(precedents[las], fir)
		}
	}

	var buffer bytes.Buffer

	var removed bool

	for {
		for i := 1; i <= 26; i++ {
			removed = false
			cha := iToStr(i)
			if !strings.Contains(buffer.String(), cha) && len(precedents[cha]) == 0 {
				buffer.WriteString(cha)
				for j, sli := range precedents {
					ind := -1
					for k := 0; k < len(sli); k++ {
						if sli[k] == cha {
							ind = k
						}
					}
					if ind != -1 {
						s := precedents[j]
						s[ind] = s[len(s)-1]
						s[len(s)-1] = ""
						s = s[:len(s)-1]
						precedents[j] = s
					}
				}
				removed = true
			}
			if removed {
				break
			}
		}
		if !removed {
			break
		}
	}

	fmt.Println(buffer.String())
}

func iToStr(i int) string {
	return string('A' - 1 + i)
}
