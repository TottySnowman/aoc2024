package main

import (
	"strconv"
	"strings"
)

func toNumbers(input string) []int {
	chars := strings.Split(input, "")
	numbers := make([]int, len(chars))
	for i, val := range chars {
		numbers[i], _ = strconv.Atoi(val)
	}
	return numbers
}
func Part1() {

	input := "2333133121414131402"

	numbers := toNumbers(input)

	totalBlocksLength := 0
	for i := 0; i < len(numbers); i += 2 { // Sum values at even indices
		totalBlocksLength += numbers[i]
	}
	blocks := make([]int, totalBlocksLength)
	currentBlockIndex := 0
	reverseIndex := len(numbers) - 1

	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			for count := 0; count < numbers[i]; count++ {
				blocks[currentBlockIndex] = i / 2
				currentBlockIndex++
			}
		} else {
			emptyFields := numbers[i]
			for count := 0; count < emptyFields; count++ {
				if currentBlockIndex >= totalBlocksLength {
					break
				}
				for numbers[reverseIndex] == 0 {
					reverseIndex -= 2
				}
				blocks[currentBlockIndex] = reverseIndex / 2
				numbers[reverseIndex]--
				currentBlockIndex++
			}
		}
	}
	checksum := 0
	for i, val := range blocks {
		checksum += i * val
	}

	println(checksum)
}
