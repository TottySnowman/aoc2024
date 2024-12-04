package main

import (
	"bufio"
	"os"
	"strings"
)

const findingText = "XMAS"

type Point struct {
	x int
	y int
}

func Part1() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var result int
	dic := make(map[int][]Point)

	lineCounter := 0
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			var points []Point
			lineCounter = lineCounter + 1
			lineStr := string(line)
			for i := 0; i < len(lineStr); i++ {
				points = append(points, Point{
					x: i,
					y: lineCounter,
				})
			}
			dic[lineCounter] = points
		}
		if err != nil {
			break
		}
	}

  

	print(result)
}

func matchWord(line string) int {
	wordCountInLine := strings.Count(line, findingText)
	reverseXMAS := reverseString(findingText)
	countTwo := strings.Count(line, reverseXMAS)
	return wordCountInLine + countTwo
}

func matchLetters(line string) {

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
