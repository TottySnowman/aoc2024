package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1() {

	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result int

	for {
		line, _, err := reader.ReadLine()
		if len(line) < 0 {
			continue
		}

		if err != nil {
			break
		}

		lineStr := string(line)
		splittedResult := strings.Split(lineStr, ":")
		targetResult, _ := strconv.Atoi(splittedResult[0])
		splittedStrings := strings.Split(splittedResult[1], " ")
		numbers := ConvertStringsToInts(splittedStrings)
		result += value(targetResult, numbers)
	}

	println(result)
}

func value(expectedResult int, numbers []int) int {
	if len(numbers) == 1 {
		if numbers[0] == expectedResult {
			return expectedResult
		}
		return 0
	}
  plusValue := numbers[0] + numbers[1] 
  multValue := numbers[0] * numbers[1] 

	tempResP := value(expectedResult, append([]int{plusValue}, numbers[2:]...))
  if tempResP != 0 {
		return tempResP
	}

	tempResM := value(expectedResult, append([]int{multValue}, numbers[2:]...))
  if tempResM != 0 {
		return tempResM
	}
	return 0
}

func ConvertStringsToInts(stringValues []string) []int {
	var splittedNumbers []int

	for _, value := range stringValues {
		num, err := strconv.Atoi(value)
		if err == nil {
			splittedNumbers = append(splittedNumbers, num)
		}
	}

	return splittedNumbers
}
