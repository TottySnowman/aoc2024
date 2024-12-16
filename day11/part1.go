package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("puzzleInput.txt")
	var stones []int
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if len(line) < 0 {
			continue
		}

		if err != nil {
			break
		}

		lineStr := string(line)
		splitted := strings.Split(lineStr, " ")
		for _, val := range splitted {
			number, _ := strconv.Atoi(val)
			stones = append(stones, number)
		}
	}

  for i := 0; i < 25; i++ {
    println("run: ", i)
    stones = blink(stones)
  }

	println(len(stones))
}

func blink(numbers []int) []int {
	var res []int
	for _, num := range numbers {

		xStr := strconv.Itoa(num)
		length := len(xStr)

		if num == 0 {
			res = append(res, 1)
		} else if len(strconv.Itoa(num))%2 == 0 {
			firstHalf, _ := strconv.Atoi(xStr[:length/2])
			secondHalf, _ := strconv.Atoi(xStr[length/2:])
			res = append(res, firstHalf, secondHalf)
		} else {
			res = append(res, (num * 2024))
		}
	}
	return res
}
