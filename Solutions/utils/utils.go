package utils

import (
	"fmt"
	"os"
)

// GetProblemInput - Returns a pointer to the file containing a
//		particular day's input
func GetProblemInput(day string) (data *os.File) {
	inputFile := "../../Inputs/Day" + day + "_Input.txt"
	data, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	return
}

// CheckError - Checks for an error and panics if err is not nil
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// PrintSolution - Prints the solution for a given puzzle
func PrintSolution(puzzle int, solution string) {
	fmt.Println("Problem", puzzle, "solution:")
	fmt.Println(solution)
	fmt.Println()
}
