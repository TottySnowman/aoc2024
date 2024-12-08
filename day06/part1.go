package main

import (
	"bufio"
	"os"
)

type Point struct {
	x int
	y int
}

type Field struct {
	hasWalked bool
	value     string
}

var result int

var roomWidth int
var linecount int
var guardPosition Point
var currentDirection string

func Part1() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	grid := map[Point]Field{}

	for {
		line, _, err := reader.ReadLine()
		if len(line) < 0 {
			continue
		}

		if err != nil {
			break
		}

		lineStr := string(line)
		roomWidth = len(lineStr)
		linecount++

		for i := 0; i < len(lineStr); i++ {

			value := string(lineStr[i])
			if string(lineStr[i]) == "^" {
				guardPosition.x = i
				guardPosition.y = linecount
				value = "."
			}
			grid[Point{i, linecount}] = Field{
				hasWalked: false,
				value:     value,
			}
		}
	}

	currentDirection = "^"
	for guardPosition.x <= roomWidth {
		moveGuard(grid)
	}

	println(result)
}

func moveGuard(grid map[Point]Field) {
	currentPos, err := grid[guardPosition]
	localDirection := currentDirection

	if !err {
		calculateWalkedThings(grid)
		println(result)
	}

	var tempPos Point = guardPosition
	if localDirection == "^" {
		tempPos.x = guardPosition.x
		tempPos.y = guardPosition.y - 1
	}

	if localDirection == ">" {

		tempPos.x = guardPosition.x + 1
		tempPos.y = guardPosition.y
	}

	if localDirection == "<" {
		tempPos.x = guardPosition.x - 1
		tempPos.y = guardPosition.y
	}

	if localDirection == "v" {
		tempPos.x = guardPosition.x
		tempPos.y = guardPosition.y + 1
	}

	tempPosVal, err := grid[tempPos]
	if !err {
		calculateWalkedThings(grid)
		println(result)
		panic("bingbong")
	}

	if tempPosVal.value == "#" {
		if localDirection == "^" {
			tempPos.x = guardPosition.x + 1
			tempPos.y = guardPosition.y
			currentDirection = ">"
		}

		if localDirection == ">" {
			tempPos.x = guardPosition.x
			tempPos.y = guardPosition.y + 1
			currentDirection = "v"
		}

		if localDirection == "<" {
			tempPos.x = guardPosition.x
			tempPos.y = guardPosition.y - 1
			currentDirection = "^"
		}

		if localDirection == "v" {
			tempPos.x = guardPosition.x - 1
			tempPos.y = guardPosition.y
			currentDirection = "<"
		}
	}


	currentPos.hasWalked = true
  grid[guardPosition] = currentPos
	guardPosition.x = tempPos.x
	guardPosition.y = tempPos.y
}

func calculateWalkedThings(grid map[Point]Field) {
	for _, value := range grid {
		if value.hasWalked {
			result++
		}
	}
}
