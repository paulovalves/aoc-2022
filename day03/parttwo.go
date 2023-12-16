package daythree

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartTwo() {
	file, err := os.Open("day03/parttwoinput")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lineOne := ""
	lineTwo := ""
	lineThree := ""
	scanner := bufio.NewScanner(file)
	currLine := 0
	var arr []rune
	var arrInts []int
	for scanner.Scan() {
		lineOne = scanner.Text()
		if scanner.Scan() {
			currLine++
			lineTwo = scanner.Text()
			if scanner.Scan() {
				currLine++
				lineThree = scanner.Text()
			}
		}
		arrTemp := compareLineChars(lineOne, lineTwo, lineThree)
		arr = append(arr, arrTemp...)
	}

	for _, l := range arr {
		arrInts = append(arrInts, int(l))
	}

	result := countAll(arrInts)
	fmt.Printf("\nDay 03 Part II: %d", result)
}

func compareLineChars(lineOne, lineTwo, lineThree string) []rune {
	var arr []rune
	for i := 0; i < len(lineOne); i++ {
		for j := 0; j < len(lineTwo); j++ {
			for k := 0; k < len(lineThree); k++ {
				if rune(lineOne[i]) == rune(lineTwo[j]) && rune(lineTwo[j]) == rune(lineThree[k]) {
					if !contains(rune(lineOne[i]), arr) {
						arr = append(arr, rune(lineOne[i]))
					}
				}
			}
		}
	}
	return arr
}
