package main

import (
	"bufio"
	"os"
	"strconv"
)

type Point struct {
	x       int
	y       int
	visited bool
}

var grid = map[Point]int{}

func Part1() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result int

	linecount := 0
	for {
		line, _, err := reader.ReadLine()
		if len(line) < 0 {
			continue
		}
		if err != nil {
			break
		}

		linecount++
		lineStr := string(line)
		for i := 0; i < len(lineStr); i++ {
			number, _ := strconv.Atoi(string(lineStr[i]))
			grid[Point{i, linecount - 1, false}] = number
		}
	}

	for i, value := range grid {
		if !isInGrid(i) {
			continue
		}
		if value != 0 {
			continue
		}
		if findIt(i) {
			result++
		}

	}

	println(result)
}

func findIt(point Point) bool {
	nbr := []Point{
		{0, -1, false}, {1, 0, false}, {0, 1, false}, {-1, 0, false},
	}

	val := grid[point]
	if val == 9 {
		return true
	}

	for _, d := range nbr {
		nextPoint := Point{x: point.x + d.x, y: point.y + d.y, visited: false}
		if !isInGrid(nextPoint) {
			continue
		}

		valNext := grid[nextPoint]
		val++
		if valNext == val {
			if findIt(nextPoint) {
				return true
			}
		}

		continue
	}

	return false
}

func isInGrid(point Point) bool {
	_, exists := grid[point]
	if !exists {
		return false
	}

	return true
}
