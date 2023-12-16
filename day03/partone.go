package daythree

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() int {
	file, err := os.Open("day03/partoneinput")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	result := 0
	var arr []rune
	var arrInts []int
	for scanner.Scan() {
		text := scanner.Text()
		half := len(text) / 2
		left := text[:half]
		right := text[half:]
		arr = compareChars(left, right)
		for _, r := range arr {
			arrInts = append(arrInts, int(r))
		}
		result = countAll(arrInts)
	}

	return result
}

func compareChars(left string, right string) []rune {
	var arr []rune
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				if !contains(rune(left[i]), arr) {
					arr = append(arr, rune(left[i]))
				}
			}
		}
	}
	return arr
}

func contains(r rune, arr []rune) bool {
	for _, v := range arr {
		if v == r {
			return true
		}
	}
	return false
}

func countAll(arr []int) int {
	result := 0
	for _, i := range arr {
		if i >= 97 {
			result += i - 96
		} else {
			result += i - 65 + 27
		}
	}
	return result
}

func DayThree() {
	result := readFile()
	fmt.Printf("\nDay 03 Part I: %d", result)
}
