package datafile

import (
	"bufio"
	"fmt"
	"github.com/DagmarC/codeOfAdvent/constants"
	"github.com/DagmarC/codeOfAdvent/task12/ship"
	"log"
	"os"
	"regexp"
	"strconv"
)

func LoadFileTask12() ([]ship.Instruction, error) {

	file, err := os.Open(constants.Task12)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var instructions []ship.Instruction
	instructions = make([]ship.Instruction, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}

		err = parseLineAndAppendInstructions(line, &instructions)
		if err != nil {
			return nil, err
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return instructions, nil
}

func parseLineAndAppendInstructions(line string, instructions *[]ship.Instruction) error {

	// line = W5
	reg := regexp.MustCompile(`^([aA-zZ]+)([0-9]+)$`)

	matches := reg.FindAllStringSubmatch(line, 1)
	if len(matches) != 1 {
		return fmt.Errorf("wrong format of: %v", line)
	}
	// matches = [[W5 W 5]]
	direction := matches[0][1]

	value, err := strconv.ParseFloat(matches[0][2], 64)
	if err != nil {
		return fmt.Errorf("wrong format %v", line)
	}
	newInstruction := ship.CreateInstruction(direction, value)
	*instructions = append(*instructions, *newInstruction)

	return nil
}
