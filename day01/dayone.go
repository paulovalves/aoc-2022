package dayone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sumArr(arr []int) int {
	result := 0
	for _, sum := range arr {
		result += sum
	}
	return result
}

func findMaxValue(arr []int) int {
	maxValue := 0
	for _, currValue := range arr {
		if currValue > maxValue {
			maxValue = currValue
		}
	}
	return maxValue
}

func DayOne() {
	fileName := "day01/day-one-input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	maxFood := 0

	var arrTemp []int
	var arrResult []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			currLine, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			arrTemp = append(arrTemp, currLine)
		} else {
			arrResult = append(arrResult, sumArr(arrTemp))
			arrTemp = nil
		}
	}

	sort.Ints(arrResult)
	lastThree := arrResult[len(arrResult)-3:]

	lt := sumArr(lastThree)

	maxFood = findMaxValue(arrResult)
	fmt.Printf("\nDay 01 Part I: %d", maxFood)
	fmt.Printf("\nDay 01 Part II: %d", lt)
}
