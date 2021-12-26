package datafile

import (
	"bufio"
	"errors"

	"log"
	"os"
	"strconv"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task8/gameBoot"
)

func LoadFileTask8() (*gameBoot.BootCode, error) {

	file, err := os.Open(constants.Task8)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	bootInstructions := gameBoot.CreateBootCode()
	bootInstructions.Initialize()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}

		err = parseLineAndSaveInstruction(line, bootInstructions)
		if err != nil {
			return nil, err
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return bootInstructions, nil
}

func parseLineAndSaveInstruction(line string, instructions *gameBoot.BootCode) error {
	// [acc, +1]
	splitLine := strings.Split(line, " ")
	if len(splitLine) != 2 {
		return errors.New("invalid line")
	}

	atomicInstr, err := gameBoot.AtomicInstructionByName(splitLine[0])
	if err != nil {
		return err
	}
	argument, err := strconv.Atoi(splitLine[1])
	if err != nil {
		return err
	}
	// If nil is not null -> create new instruction.
	instruction, err := instructions.InstructionByNameArg(atomicInstr, argument)
	if err != nil {
		instruction = gameBoot.CreateInstruction(atomicInstr, argument)
	}
	*instructions = append(*instructions, instruction)

	return nil
}
