package dayfour

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartTwo() {
	file, err := os.Open("day04/part-two-input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()

		left, right := splitLine(text, ",")
		n1, n2, n3, n4 := makeNumbers(left, right)
		result += checkNumbers(n1, n2, n3, n4)
	}
	fmt.Println(result)
}

func checkNumbers(num1, num2, num3, num4 int) int {
	if num1 >= num3 && num1 <= num4 {
		return 1
	} else if num1 <= num3 && num2 >= num3 {
		return 1
	} else if num1 <= num4 && num2 >= num4 {
		return 1
	} else {
		return 0
	}
}
