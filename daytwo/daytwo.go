package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	ROCK     string
	PAPER    string
	SCISSORS string
}

func calculatePoints() {

	fileName := "daytwo/input"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		el := scanner.Text()
		l, r := splitElement(el)
		fmt.Printf("\n%s -> %s", l, r)
	}

	/*for _, el := range arr {
		l, r := splitElement(el)
		fmt.Printf("\n%s -> %s", l, r)
	}*/
}

func splitElement(str string) (string, string) {
	s := strings.Split(str, " ")
	left := s[0]
	right := s[1]
	return left, right
}

func DayTwo() {
	calculatePoints()
}
