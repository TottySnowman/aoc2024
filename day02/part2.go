package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var validFloors int

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			lineStr := string(line)
			res := strings.Split(lineStr, " ")
			var numbers []int
			for i := 0; i < len(res); i++ {
				number, _ := strconv.Atoi(res[i])
				numbers = append(numbers, number)
			}

			if isReportSafe(numbers) {
				validFloors++
			}
		}
		if err != nil {
			break
		}
	}

	print(validFloors)
}

func isReportSafeWithDamper(reportNum []int) bool {
	flagIncrease, flagDecrease, damperUsed := false, false, false
  

	for i := 1; i < len(reportNum); i++ {
		diff := reportNum[i] - reportNum[i-1]

		if diff > 0 {
			flagIncrease = true
		} else if diff < 0 {
			flagDecrease = true
		} else if !damperUsed {
			damperUsed = true
		} else {
			return false
		}

		if flagDecrease && flagIncrease && damperUsed {
			return false
		}
		if flagDecrease && flagIncrease {
			damperUsed = true
		}

		if (diff > 3 || diff < -3) && damperUsed {
			return false
		}

		if diff > 3 || diff < -3 {
			damperUsed = true
		}
	}

	return true
}
