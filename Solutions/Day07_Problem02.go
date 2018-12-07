package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"./Packages/errors"
)

type worker struct {
	task  string
	start int
	end   int
}

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
	var workers [5]worker
	var available []string

	for sec := 0; sec < 2400; sec++ {
		for i, w := range workers {
			if w.end <= sec {
				buffer.WriteString(w.task)
				workers[i].task = ""
				completeTask(w.task, precedents)
				available = getAvailableTasks(buffer, precedents)
				available = reduceAvailable(available, workers)
				sort.Strings(available)
				if len(available) > 0 {
					workers[i].task = available[0]
					workers[i].start = sec
					workers[i].end = int([]rune(workers[i].task)[0]) - 4 + sec
				}
			}
		}
		if len(buffer.String()) == 26 {
			fmt.Println(sec - 1)
			break
		}
	}
}

func iToStr(i int) string {
	return string('A' - 1 + i)
}

func getAvailableTasks(finished bytes.Buffer, p map[string][]string) (result []string) {
	for i := 1; i <= 26; i++ {
		cha := iToStr(i)
		if !strings.Contains(finished.String(), cha) && len(p[cha]) == 0 {
			result = append(result, cha)
		}
	}
	return
}

func reduceAvailable(tasks []string, w [5]worker) []string {
	result := make([]string, len(tasks))
	copy(result, tasks)
	for _, work := range w {
		cha := work.task
		for i := 0; i < len(result); i++ {
			if result[i] == cha {
				result = append(result[:i], result[i+1:]...)
			}
		}
	}
	return result
}

func completeTask(task string, p map[string][]string) {
	for j, sli := range p {
		ind := -1
		for k := 0; k < len(sli); k++ {
			if sli[k] == task {
				ind = k
			}
		}
		if ind != -1 {
			s := p[j]
			s[ind] = s[len(s)-1]
			s[len(s)-1] = ""
			s = s[:len(s)-1]
			p[j] = s
		}
	}
}
