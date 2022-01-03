package datafile

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/utils"
)

// LoadFileTask6 parses task6 file. Each line represents answers from 1 person. Empty line separates groups.
// Group - represented via keys in map and person is represented via value in group. Group:person pair.
// FILE:
//whley
//dhyetlqf
//
//fqx
//xf
//bfxz
func LoadFileTask6() (*map[int]string, *map[int][]string, error) {

	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		return nil, nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	allUniqueAnswers := make(map[int]string)
	allAnswers := make(map[int][]string)
	groupNumber := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		// Load next group.
		if line == "" {
			groupNumber++
			continue
		}
		// TASK 1 - Anyone answers yes to question.
		addUniqueAnswersToGroup(line, &allUniqueAnswers, groupNumber)
		// TASK 2 - All in 1 group answers yes to question.
		addAllAnswersToGroup(line, &allAnswers, groupNumber)
	}
	if scanner.Err() != nil {
		return nil, nil, err
	}

	return &allUniqueAnswers, &allAnswers, nil
}

func addAllAnswersToGroup(answer string, allAnswers *map[int][]string, groupNumber int) {
	(*allAnswers)[groupNumber] = append((*allAnswers)[groupNumber], answer)
}

// addUniqueAnswersToGroup identify the questions to which anyone answered "yes";
func addUniqueAnswersToGroup(personAnswers string, allAnswers *map[int]string, groupNumber int) {
	for _, a := range personAnswers {
		if !strings.Contains((*allAnswers)[groupNumber], string(a)) {
			(*allAnswers)[groupNumber] += string(a)
		}
	}
}
