package dayfive

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

func PartOne() {
	arr, moves, err := parseInputCrates()
	if err != nil {
		fmt.Println(err)
	}
	r := arrange(arr)

	result := moveCrates(r, moves)
	final := getTop(result)
	fmt.Printf("\nDay 05 Part I: %v", final)
	// fmt.Println(final)
}

func getTop(arr [][]string) []string {
	var finalArr []string

	for i := 0; i < len(arr); i++ {
		finalArr = append(finalArr, arr[i][len(arr[i])-1])
	}

	return finalArr
}

func arrange(arr [][]string) [][]string {
	var tempArr []string
	var finalArr [][]string
	arr[0][len(arr[0])-1] = "   "

	// fmt.Printf("\nArr: %d\n", len(arr))
	for i := 0; i < len(arr); i++ {
		for j := len(arr[i]) - 2; j > -1; j-- {
			if strings.TrimSpace(arr[j][i]) != strings.TrimSpace("") {
				tempArr = append(tempArr, arr[j][i])
			}
		}
		// var a []string
		//
		// if len(tempArr) > 0 {
		// 	b := len(tempArr)
		// 	for i := 0; i < len(tempArr); i++ {
		// 		a = append(a, tempArr[b-1])
		// 		b--
		// 	}
		// }
		finalArr = append(finalArr, tempArr)
		tempArr = nil
	}
	tempArr = nil
	for i := len(arr) - 1; i > -1; i-- {
		if strings.TrimSpace(arr[i][len(arr)]) != "" {
			tempArr = append(tempArr, strings.TrimSpace(arr[i][len(arr)]))
		}
	}
	if tempArr != nil {
		finalArr = append(finalArr, tempArr)
	}

	return finalArr
}

func moveCrates(arr [][]string, moves [][]int) [][]string {
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

		result = newMove(result, q, c, d)

	}

	return result
}

func newMove(arr [][]string, quantity, stack, destination int) [][]string {
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

	for i := 0; i < len(crates); i++ {
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

func makeMoves() [][]int {
	moves := `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	var arr [][]int
	var tempArr []int
	words := strings.Fields(moves)
	for i := 1; i < len(words); i += 2 {
		w, _ := strconv.Atoi(words[i])
		tempArr = append(tempArr, w)
		if len(tempArr) == 3 {
			arr = append(arr, tempArr)
			tempArr = nil
		}
	}
	return arr
}

func parseMoves(line string) []int {
	var arr []int
	words := strings.Fields(line)
	for i := 1; i < len(words); i += 2 {
		w, _ := strconv.Atoi(words[i])
		arr = append(arr, w)
	}
	return arr
}

func parseInputCrates() ([][]string, [][]int, error) {
	file, err := os.Open("day05/part-one-input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr [][]string
	var moves [][]int
	var text string

	for scanner.Scan() {
		text = scanner.Text()
		if strings.Contains(text, "move") {
			moves = append(moves, parseMoves(text))
		} else if len(text) < 1 {
			continue
		} else if text == " 1   2   3   4   5   6   7   8   9 " {
			continue
		} else {
			arr = append(arr, makeArray(text))
		}

	}
	return arr, moves, nil
}

func makeArray(text string) []string {
	var tempArr []string
	var line string
	var scan *bufio.Reader

	for {
		scan = bufio.NewReader(strings.NewReader(text))
		for {
			char, _, _ := scan.ReadRune()
			if char == 0 {
				if len(strings.TrimSpace(line)) == 0 {
					line = "   "
					tempArr = append(tempArr, strings.TrimSpace(line))
					return tempArr
				} else {
					tempArr = append(tempArr, strings.TrimSpace(line))
					return tempArr
				}
				// line = ""
			} else if len(line) < 3 {
				line = fmt.Sprintf("%s%c", line, char)
			} else if len(line) == 3 && len(strings.TrimSpace(line)) == 0 {
				line = "   "
				tempArr = append(tempArr, line)
				line = ""
			} else if len(line) == 3 && len(strings.TrimSpace(line)) != 0 {
				tempArr = append(tempArr, strings.TrimSpace(line))
				line = ""
			} else {
				fmt.Println("else")
			}
		}
	}
}

func parseFile() ([][]string, error) {
	str := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
 `

	var lineArr [][]string
	var tempArr []string
	scn := bufio.NewReader(strings.NewReader(str))
	line := ""
	for {
		char, _, err := scn.ReadRune()
		if err == io.EOF {
			lineArr = lineArr[:len(lineArr)-1]
			return lineArr, nil
		}
		// fmt.Println(char)
		if string(char) == "\n" {
			if len(strings.TrimSpace(line)) == 0 {
				line = "   "
				tempArr = append(tempArr, line)
			} else {
				// fmt.Printf("\nn %s", line)
				tempArr = append(tempArr, strings.TrimSpace(line))
			}
			lineArr = append(lineArr, tempArr)
			line = ""
			tempArr = nil
		} else if len(line) < 3 {
			line = fmt.Sprintf("%s%c", line, char)
		} else if len(line) == 3 && len(strings.TrimSpace(line)) == 0 {
			line = "   "
			tempArr = append(tempArr, line)
			// fmt.Printf("\n0 %s", line)
			line = ""
		} else if len(line) == 3 && len(strings.TrimSpace(line)) != 0 {
			tempArr = append(tempArr, strings.TrimSpace(line))
			// fmt.Printf("\n1 %s", line)
			line = ""

		} else {
			break
		}

	}
	lineArr = lineArr[:len(lineArr)-1]
	return lineArr, nil
}
