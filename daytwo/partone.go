package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Player: A for Rock, B for Paper, and C for Scissors
// Points: 1 for Rock, 2 for Paper, and 3 for Scissors
// Player: X for Rock, Y for Paper, and Z for Scissors
// outcome: 0 - lost, 3 - draw, 6 - win

type Game struct {
	ROCK     string
	PAPER    string
	SCISSORS string
}

func calculatePoints() int {
	fileName := "daytwo/inputpartone"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		el := scanner.Text()
		l, r := splitElement(el)
		_, n := outcomePartOne(l, r)
		sum += n
		// fmt.Printf("\n%s = %d", s, n)
	}

	/*for _, el := range arr {
		l, r := splitElement(el)
		fmt.Printf("\n%s -> %s", l, r)
	}*/
	return sum
}

func outcomePartOne(left string, right string) (string, int) {
	// Rock | Rock
	if left == "A" && right == "X" {
		return "draw", 1 + 3
		// rock | paper
	} else if left == "A" && right == "Y" {
		return "WON", 2 + 6
		// rock | scissors
	} else if left == "A" && right == "Z" {
		return "Lost", 3 + 0
		// paper | rock
	} else if left == "B" && right == "X" {
		return "Lost", 1 + 0
		// paper | paper
	} else if left == "B" && right == "Y" {
		return "Draw", 2 + 3
		// paper | scissors
	} else if left == "B" && right == "Z" {
		return "WON", 3 + 6
		// scissors | rock
	} else if left == "C" && right == "X" {
		return "WON", 1 + 6
		// scissors | paper
	} else if left == "C" && right == "Y" {
		return "Lost", 2 + 0
		// scissors | scissors
	} else if left == "C" && right == "Z" {
		return "Draw", 3 + 3
	} else {
		return "WRONG", 0
	}
}

func splitElement(str string) (string, string) {
	s := strings.Split(str, " ")
	left := s[0]
	right := s[1]
	return left, right
}

func PartOne() {
	result := calculatePoints()
	fmt.Printf("\n\nResult: %d\n", result)
}
