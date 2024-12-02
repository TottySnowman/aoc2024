package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

  reader := bufio.NewReader(file)

	listOne, listTwo := GetLists(reader)

	listOne = SortListsByNumberAsc(listOne)
	listTwo = SortListsByNumberAsc(listTwo)
	var counter int

	for i := 0; i < len(listOne); i++ {
		distance := listOne[i] - listTwo[i]
		if distance < 0 {
			distance = -distance
		}

		counter = counter + distance
	}
	print(counter)
}

func GetLists(reader *bufio.Reader) ([]int, []int) {
	var listOne []int
	var listTwo []int

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			lineStr := string(line)
			res := strings.Split(lineStr, "   ")
			v1, _ := strconv.Atoi(res[0])
			v2, _ := strconv.Atoi(res[1])
			listOne = append(listOne, v1)
			listTwo = append(listTwo, v2)
		}
		if err != nil {
			break
		}
	}

	return listOne, listTwo
}

func SortListsByNumberAsc(inputList []int) []int {
	sort.Slice(inputList, func(i, j int) bool {
		return inputList[i] < inputList[j]
	})

	return inputList
}
