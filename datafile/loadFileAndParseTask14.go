package datafile

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/DagmarC/codeOfAdvent/constants"
	"github.com/DagmarC/codeOfAdvent/task14/initprogram"
)

func LoadFileAndRunTask14(subtask int) (*initprogram.Instructions, error) {
	file, err := os.Open(constants.Task14)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	instructions := initprogram.InitializeInitProgram()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		err := parseLineOfProgram(line, instructions, subtask)
		if err != nil {
			return nil, err
		}

	}
	if scanner.Err() != nil {
		return nil, err
	}

	return instructions, nil
}

func parseLineOfProgram(line string, instructions *initprogram.Instructions, subtask int) error {
	if ok, _ := regexp.MatchString("mask = ", line); ok {
		bitmask, err := parseMask(line)
		if err != nil {
			return err
		}
		instructions.UpdateBitmask(bitmask)

	} else {
		memLocation, value, err := parseMemLocation(line)
		if err != nil {
			return err
		}
		memUint := parseToUint64(memLocation)
		valueUint := parseToUint64(value)
		if subtask == 1 {
			instructions.Write(memUint, valueUint)
		} else if subtask == 2 {
			instructions.WriteSubtask2(memUint, valueUint)

		}

	}
	return nil
}

func parseToUint64(toParse string) uint64 {
	parsed, err := strconv.ParseUint(toParse, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}

func parseMemLocation(line string) (string, string, error) {

	pattern := `^mem\[(\d+)\] = (\d+)$`
	reg := regexp.MustCompile(pattern)

	matches := reg.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 && len(matches[0]) != 3 {
		return "nil", "nil", errors.New("invalid line")
	}
	memLocation := matches[0][1]
	value := matches[0][2]

	return memLocation, value, nil

}

func parseMask(line string) (string, error) {

	pattern := `^mask = ([X01]{36})$`
	reg := regexp.MustCompile(pattern)

	matches := reg.FindAllStringSubmatch(line, -1)
	return matches[0][1], nil
}
