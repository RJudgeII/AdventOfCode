package main

import (
	"bufio"
	"bytes"
	"sort"
	"strconv"
	"strings"

	"../utils"
)

type worker struct {
	task  string
	start int
	end   int
}

func main() {
	precedents := setMap()

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

	utils.PrintSolution(1, buffer.String())

	precedents2 := setMap()

	var buffer2 bytes.Buffer
	var workers [5]worker
	var available []string
	var seconds string

	for sec := 0; sec < 2400; sec++ {
		for i, w := range workers {
			if w.end <= sec {
				buffer2.WriteString(w.task)
				workers[i].task = ""
				completeTask(w.task, precedents2)
				available = getAvailableTasks(buffer2, precedents2)
				available = reduceAvailable(available, workers)
				sort.Strings(available)
				if len(available) > 0 {
					workers[i].task = available[0]
					workers[i].start = sec
					workers[i].end = int([]rune(workers[i].task)[0]) - 4 + sec
				}
			}
		}
		if len(buffer2.String()) == 26 {
			seconds = strconv.Itoa(sec - 1)
			break
		}
	}

	utils.PrintSolution(2, seconds)
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

func setMap() map[string][]string {
	data := utils.GetProblemInput("07")
	defer data.Close()

	output := make(map[string][]string)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		fir := s[1]
		las := s[7]

		if len(output[las]) == 0 {
			output[las] = []string{fir}
		} else {
			output[las] = append(output[las], fir)
		}
	}

	return output
}
