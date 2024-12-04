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

	grid := map[Point]string{}

	linecount := 0
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			linecount++
			lineStr := string(line)
			for i := 0; i < len(lineStr); i++ {
				grid[Point{i, linecount}] = string(lineStr[i])
			}
		}
		if err != nil {
			break
		}
	}

	result := 0
	for p := range grid {
		result += strings.Count(strings.Join(findMatchingWord(p, grid), " "), findingText)
	}
	print(result)
}
func findMatchingWord(point Point, grid map[Point]string) []string {
	delta := []Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	}

	words := make([]string, len(delta))
	for _, d := range delta {
		word := ""
		for n := 0; n < 4; n++ {
			nextPoint := Point{x: point.x + d.x*n, y: point.y + d.y*n}
			word += grid[nextPoint]
		}
		words = append(words, word)
	}

	return words
}
