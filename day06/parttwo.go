package daysix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() {
	file, err := os.Open("day06/part-one-input")
	if err != nil {
		log.Fatal(err)
	}
	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	pos, result := checkString(text)
	fmt.Printf("\nDay 06 Part II: %s - %d", result, pos)
}

func checkString(text string) (int, string) {
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
		if len(line) < 14 {
			char, _, _ := scan.ReadRune()
			line = fmt.Sprintf("%s%c", line, char)
		} else {
			if len(line) == 14 {
				result, unique = checkUniques(line)
				if unique {
					position = countPosition(textToCheck, text)
					return position, result
				}
			}
			if len(textToCheck) < 13 {
				break
			}
			counter++
			textToCheck = text[counter-1:]

			line = ""

		}
	}
	return position, result
}
