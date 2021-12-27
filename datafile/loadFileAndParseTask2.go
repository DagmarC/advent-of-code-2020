// Package datafile of Task2 will parse file input: 3-4 l: vdcv and saves it into passp.Definition structure.
package datafile

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task2/passp"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

// LoadFileTask2 loads, parses creates and returns []*passp.Definition
func LoadFileTask2() ([]*passp.Definition, error) {

	file, err := os.Open(path.Join(utils.GetWd(), constants.Input))
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var passwordDefinitions []*passp.Definition

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Save parsed line into the struct.
		pd, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		passwordDefinitions = append(passwordDefinitions, pd)
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return passwordDefinitions, nil
}

func parseLine(line string) (*passp.Definition, error) {
	// result: [3-4, l:, vdcv]
	dataSplit := strings.Split(line, " ")
	if len(dataSplit) != 3 {
		return nil, errors.New("data in file corrupted")
	}

	// result: [3, 4]
	rangeSplit := strings.Split(dataSplit[0], "-")
	lowerBoundary, err := strconv.Atoi(rangeSplit[0])
	if err != nil {
		return nil, err
	}
	upperBoundary, err := strconv.Atoi(rangeSplit[1])
	if err != nil {
		return nil, err
	}

	// result: [l]
	letter := strings.Split(dataSplit[1], ":")[0]

	// result: [vdcv]
	password := dataSplit[2]

	pp := passp.CreatePasswordPolicy(letter, lowerBoundary, upperBoundary)
	passwordDefinition := passp.CreatePasswordDefinition(password, pp)

	return passwordDefinition, nil
}
