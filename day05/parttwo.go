package dayfive

import (
	"fmt"
	"os"
)

func PartTwo() {
	// arr, err := parseFile()
	arr, moves, err := parseInputCrates()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(arr)
	r := arrange(arr)

	// fmt.Println(r)
	// fmt.Println(moves)
	result := moveCratesPartTwo(r, moves)
	// fmt.Println(result)
	final := getTop(result)
	fmt.Printf("\nDay 05 Part II: %v", final)
}

func moveCratesPartTwo(arr [][]string, moves [][]int) [][]string {
	result := arr
	for i := 0; i < len(moves); i++ {
		// for i := 0; i < 40; i++ {
		j := 0
		q := moves[i][j]
		c := moves[i][j+1] - 1
		d := moves[i][j+2] - 1
		if c < 0 {
			c = 0
		}
		if d < 0 {
			d = 0
		}

		result = moveSameOrder(result, q, c, d)

	}

	return result
}

func moveSameOrder(arr [][]string, quantity, stack, destination int) [][]string {
	var crates []string
	var maybe []string
	crate := len(arr[stack]) - 1

	for i := quantity; i > 0; i-- {

		if crate == -1 {
			break
		}
		crates = append(crates, arr[stack][crate])
		crate--
	}

	for i := len(crates) - 1; i > -1; i-- {
		arr[destination] = append(arr[destination], crates[i])
	}
	if len(crates) == len(arr[stack]) {
		arr[stack] = maybe
	} else if len(crates) < len(arr[stack]) {
		r := len(arr[stack]) - len(crates)
		arr[stack] = arr[stack][:r]
	} else {
		arr[stack] = arr[stack][quantity:]
	}

	file, err := os.OpenFile("day05/yourfile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return arr
	}
	defer file.Close()
	content := fmt.Sprintf("%v\n", arr)

	// Write the formatted content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return arr
	}

	return arr
}
