package main

import (
	"bufio"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			lineStr := string(line)
			match := re.FindAllStringSubmatch(lineStr, -1)

      for _, matching := range match {
        cleanedString := strings.ReplaceAll(strings.ReplaceAll(matching[0], "mul(", ""), ")", "")
        splitted := strings.Split(cleanedString, ",")
        number1, _ := strconv.Atoi(splitted[0])
        number2, _ := strconv.Atoi(splitted[1])
        res := number1 * number2

        result = result + res
      }
		}
		if err != nil {
			break
		}
	}
	print(result)
}
