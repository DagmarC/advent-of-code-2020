package datafile

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/DagmarC/advent-of-code-2020/utils"
)

func LoadFileTask9() ([]int, error) {

	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	numbers := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}

		err = parseLineAndAppendNumbers(line, &numbers)
		if err != nil {
			return nil, err
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return numbers, nil
}

func parseLineAndAppendNumbers(line string, numbers *[]int) error {
	n, err := strconv.Atoi(line)
	if err != nil {
		return errors.New("NaN")
	}
	*numbers = append(*numbers, n)
	return nil
}
