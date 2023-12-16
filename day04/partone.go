package dayfour

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	readFile()
	return 0
}

func readFile() {
	file, err := os.Open("day04/part-one-input")
	if err != nil {
		log.Fatal(err)
	}

	var result int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		left, right := splitLine(text, ",")
		num1, num2, num3, num4 := makeNumbers(left, right)
		result += checkRange(num1, num2, num3, num4, result)
	}

	fmt.Println(result)
}

func splitLine(str, sep string) (string, string) {
	split := strings.Split(str, sep)
	left := split[0]
	right := split[1]

	return left, right
}

func makeNumbers(left, right string) (int, int, int, int) {
	ll, lr := splitLine(left, "-")
	rl, rr := splitLine(right, "-")
	num1, _ := strconv.Atoi(ll)
	num2, _ := strconv.Atoi(lr)
	num3, _ := strconv.Atoi(rl)
	num4, _ := strconv.Atoi(rr)

	return num1, num2, num3, num4
}

func checkRange(n1, n2, n3, n4, result int) int {
	var n int
	if n1 == n2 {
		n = n1 + n2
		n = n / 2
		if n >= n3 && n <= n4 {
			return 1
		}
	}
	if n1 >= n3 && n2 <= n4 {
		return 1
	}

	if n3 >= n1 && n4 <= n2 {
		return 1
	}

	return 0
}
