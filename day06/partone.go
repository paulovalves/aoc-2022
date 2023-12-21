package daysix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartOne() {
	file, err := os.Open("day06/part-one-input")
	if err != nil {
		log.Fatal(err)
	}
	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	pos, result := checkStringSlice(text)
	fmt.Printf("\nDay 06 Part I: %s - %d", result, pos)
}

func checkStringSlice(text string) (int, string) {
	textToCheck := text
	counter := 0
	line := ""
	unique := false
	position := 0
	result := ""

	scan := bufio.NewReader(strings.NewReader(textToCheck))
	for {
		if len(line) == 0 {
			scan = bufio.NewReader(strings.NewReader(textToCheck))
		}
		if len(line) < 4 {
			char, _, _ := scan.ReadRune()
			line = fmt.Sprintf("%s%c", line, char)
		} else {
			if len(line) == 4 {
				result, unique = checkUniques(line)
				if unique {
					position = countPosition(textToCheck, text)
					return position, result
				}
			}
			if len(textToCheck) < 3 {
				break
			}
			counter++
			textToCheck = text[counter-1:]

			line = ""

		}
	}
	return position, result
}

func checkUniques(slice string) (string, bool) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				return "", false
			}
		}
	}
	return slice, true
}

func countPosition(slice, text string) int {
	return len(text) - len(slice) + 4
}
