package datafile

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/utils"
)

func ReadFile(input string) string {
	fileBuffer, err := ioutil.ReadFile(input)
	utils.Check(err)

	return strings.TrimSpace(string(fileBuffer))
}

func ReadLines(input string) []string {

	file, err := os.Open(input)
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	utils.Check(scanner.Err())

	return lines
}

// ReadLineOfNumbers input type "11,0,1,10,5,19"
func ReadLineOfNumbers(input string) []int {

	file, err := os.Open(input)
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := make([]int, 0)

	for scanner.Scan() {
		numbersStr := strings.Split(scanner.Text(), ",")
		for _, number := range numbersStr {
			numbers = append(numbers, utils.ToInt(number))
		}
	}
	utils.Check(scanner.Err())

	return numbers
}
