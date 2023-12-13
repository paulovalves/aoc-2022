package daythree

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() int {
	file, err := os.Open("daythree/input")
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
	fmt.Printf("\nleft: %s - right: %s", left, right)
	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				if !contains(rune(left[i]), arr) {
					fmt.Printf("\n%s == %s", string(left[i]), string(right[j]))
					arr = append(arr, rune(left[i]))

				}
			}
		}
	}
	// fmt.Println(h)
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
		// fmt.Println(i)
		if i >= 97 {
			result += i - 96
		} else {
			result += i - 65 + 27
		}
	}
	return result
}

func countChars(arr []rune) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		n := int(arr[i])
		if n >= 97 {
			result += n - 96
		} else {
			result += n - 65 + 27
		}
	}
	return result
}

func DayThree() {
	result := readFile()
	fmt.Printf("\n\nResultado: %d\n", result)
}
