package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var numbersOnly []int

func Part1() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result int
	var input string

	var fileSystemStringy string
	for {
		line, _, err := reader.ReadLine()
		if len(line) < 0 {
			continue
		}

		if err != nil {
			break
		}
		input = string(line)
	}

	isSpace := false
	index := 0

	for i := 0; i < len(input); i++ {
		number, _ := strconv.Atoi(string(input[i]))

		if isSpace {
			for j := 0; j < number; j++ {
				fileSystemStringy = fileSystemStringy + "."
			}
		} else {
			for j := 0; j < number; j++ {
				numbersOnly = append(numbersOnly, index)
				fileSystemStringy = fileSystemStringy + strconv.Itoa(index)
			}
			index++
		}
		isSpace = !isSpace
	}

  for strings.Contains(fileSystemStringy, "."){
    fileSystemStringy = replaceSingleDot(fileSystemStringy)
  }
  println(fileSystemStringy)
	println(result)
}

func replaceSingleDot(input string) string {
	number := numbersOnly[len(numbersOnly)-1]
  numbersOnly = numbersOnly[:len(numbersOnly)-1]
  println(numbersOnly)
	inputBytes := []byte(input)
  
	for i := 0; i < len(inputBytes); i++ {
		if inputBytes[i] == '.' {
			inputBytes[i] = byte('0' + number)
      break
		}
	}

  inputBytes = inputBytes[:len(inputBytes)-1]
	return string(inputBytes)
}
