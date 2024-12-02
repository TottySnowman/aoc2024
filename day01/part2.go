package main

import (
	"bufio"
	"os"
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
  listOne, listTwo := GetListsP2(reader)

  var resultCount int
  for _, value := range listOne {
    count, exists := listTwo[value]

    if exists{
      resultCount = resultCount + (count * value)
    }
  }

  print(resultCount)
}

func GetListsP2(reader *bufio.Reader)([]int, map[int]int){
	var listOne []int
	var listTwo = make(map[int]int)

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			lineStr := string(line)
			res := strings.Split(lineStr, "   ")
			v1, _ := strconv.Atoi(res[0])
			v2, _ := strconv.Atoi(res[1])
      listOne = append(listOne, v1)

      value, exists := listTwo[v2]

      if exists{
        value = value + 1
        listTwo[v2] = value
      }else{
        listTwo[v2] = 1
      }
		}
		if err != nil {
			break
		}
	}

  return listOne, listTwo
}
