// https://adventofcode.com/2020/day/8#part1
// https://adventofcode.com/2020/day/8#part2
// acc increases or decreases a single global value called the accumulator by the value given in the argument.
//For example, acc +7
// jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument
//as an offset from the jmp instruction;
// nop stands for No OPeration - it does nothing.
// This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction a second time, you know it will never terminate.
//
//Immediately before the program would run an instruction a second time, the value in the accumulator is 5.
//
//Run your copy of the boot code.
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/task8/gameBoot"
)

func main() {
	instructionsCode, err := datafile.LoadFileTask8()
	if err != nil {
		log.Fatal(err)
	}
	// TASK 1
	_, _ = readCode(instructionsCode)
	fmt.Println("Acc global value after 1st run before cycle appeared:", gameBoot.AccValue)

	// TASK 2
	gameBoot.AccValue = 0
	findCodeMistake(*instructionsCode)
	fmt.Println("Acc global value after cycle is removed is:", gameBoot.AccValue)

}

// TASK 1 + part of 2
func readCode(bootCode *gameBoot.BootCode) (map[int]*gameBoot.Instruction, error) {
	index := 0
	var codeSteps map[int]*gameBoot.Instruction
	codeSteps = make(map[int]*gameBoot.Instruction)

	for {
		if index >= len(*bootCode) {
			break
		}
		if _, ok := codeSteps[index]; ok {
			return codeSteps, errors.New("cycle appeared")
		}
		currentInstruction := (*bootCode)[index]
		if currentInstruction.AInstruction() != gameBoot.Acc {
			// Save the step at current index (not the acc instruction).
			//For task 2 it is more effective to save only nop and jmp.
			codeSteps[index] = currentInstruction
		}
		performInstruction(&index, (*bootCode)[index])
	}
	return codeSteps, nil
}

func performInstruction(index *int, instruction *gameBoot.Instruction) {
	switch instruction.AInstruction() {
	case gameBoot.Acc:
		gameBoot.AccValue += instruction.Argument()
		*index++
	case gameBoot.Jmp:
		*index += instruction.Argument()
	case gameBoot.Nop:
		*index++
	}
}

func switchInstructionJmpNop(instr *gameBoot.Instruction) {
	switch instr.AInstruction() {
	case gameBoot.Jmp:
		instr.SetAInstruction(gameBoot.Nop)
	case gameBoot.Nop:
		instr.SetAInstruction(gameBoot.Jmp)
	}
}

// TASK 2
// Fix the program so that it terminates normally by changing exactly one jmp (to nop) or nop (to jmp).
// What is the value of the accumulator after the program terminates?
func findCodeMistake(bootCode gameBoot.BootCode) {

	// Read the code 1st time to obtain all the steps that could cause harm and then change only those steps.
	codeSteps, err := readCode(&bootCode)
	if err == nil { // Err is not nil only if the cycle has been appeared.
		return
	}
	// Loop over visited codeSteps and make the changes until the code terminates.
	for _, visitedInstruction := range codeSteps {
		// Reset the acc counter.
		gameBoot.AccValue = 0

		// Change 1 by 1 already visited instructions from nop to jmp OR jmp to nop until the code terminates successfully.
		switchInstructionJmpNop(visitedInstruction)

		if _, err := readCode(&bootCode); err == nil {
			fmt.Println("TERMINATION HAS BEEN SUCCESSFUL") // Code terminated successfully.
			break
		}
		// Switch instruction back, if not successful.
		switchInstructionJmpNop(visitedInstruction)
	}
}
