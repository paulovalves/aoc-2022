package daythree

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
// Player: A for Rock, B for Paper, and C for Scissors
// Points: 1 for Rock, 2 for Paper, and 3 for Scissors
// Player: X for Rock, Y for Paper, and Z for Scissors
// outcome: 0 - lost, 3 - draw, 6 - win

func readFile() int {
	fileName := "daythree/input"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		l, r := splitString(text)
		_, n := outcome(l, r)
		sum += n
	}
	return sum
}

func splitString(str string) (string, string) {
	spt := strings.Split(str, " ")
	l := spt[0]
	r := spt[1]
	return l, r
}

func outcome(left string, right string) (string, int) {
	if left == "A" && right == "X" {
		return "Lost", 3
	} else if left == "A" && right == "Y" {
		return "Draw", 4
	} else if left == "A" && right == "Z" {
		return "WON", 8
	} else if left == "B" && right == "X" {
		return "Lost", 1
	} else if left == "B" && right == "Y" {
		return "Draw", 5
	} else if left == "B" && right == "Z" {
		return "WON", 9
	} else if left == "C" && right == "X" {
		return "Lost", 2
	} else if left == "C" && right == "Y" {
		return "Draw", 6
	} else if left == "C" && right == "Z" {
		return "WON", 7
	} else {
		return "Wrong", 0
	}
}

func DayThree() {
	result := readFile()
	fmt.Printf("\n\n%d\n", result)
}
