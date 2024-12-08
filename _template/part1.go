package main

import (
	"bufio"
	"os"
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
    println(lineStr)
	}

	println(result)
}
