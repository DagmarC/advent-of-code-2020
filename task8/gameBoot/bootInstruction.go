package gameBoot

import (
	"errors"
	"fmt"
)

type AtomicInstruction string

const (
	Nop AtomicInstruction = "nop"
	Acc AtomicInstruction = "acc"
	Jmp AtomicInstruction = "jmp"
)

var AccValue = 0

type Instruction struct {
	aInstruction AtomicInstruction
	argument     int
}

func (a *AtomicInstruction) String() string {
	switch *a {
	case Nop:
		return "nop"
	case Acc:
		return "acc"
	case Jmp:
		return "jmp"
	}
	return ""
}

func CreateInstruction(instruction AtomicInstruction, argument int) *Instruction {
	return &Instruction{
		aInstruction: instruction,
		argument:     argument,
	}
}

func (i *Instruction) Equals(otherInstruction *Instruction) bool {
	return i.aInstruction == otherInstruction.aInstruction &&
		i.argument == otherInstruction.argument
}

func (i *Instruction) AInstruction() AtomicInstruction {
	return i.aInstruction
}

func (i *Instruction) SetAInstruction(aInstruction AtomicInstruction) {
	i.aInstruction = aInstruction
}

func (i *Instruction) Argument() int {
	return i.argument
}

func (i *Instruction) SetArgument(argument int) {
	i.argument = argument
}

func AtomicInstructionByName(name string) (AtomicInstruction, error) {
	switch name {
	case "acc":
		return Acc, nil
	case "nop":
		return Nop, nil
	case "jmp":
		return Jmp, nil
	}
	return AtomicInstruction(""), errors.New("no instruction match")
}

// *********************** INSTRUCTIONS [] *************************

type BootCode []*Instruction

func CreateBootCode() *BootCode {
	return &BootCode{}
}

func (b *BootCode) Initialize() {
	*b = make(BootCode, 0)
}

func (b *BootCode) AddInstruction(instruction *Instruction) {
	*b = append(*b, instruction)
}

func (b *BootCode) InstructionByNameArg(name AtomicInstruction, argument int) (*Instruction, error) {
	for _, currInstruction := range *b {
		if currInstruction.aInstruction == name && currInstruction.argument == argument {
			return currInstruction, nil
		}
	}
	return nil, errors.New("instruction not found")
}

func (b *BootCode) PrintInstructions() {
	for _, currInstruction := range *b {
		fmt.Println(currInstruction)
	}
}

func (b *BootCode) Exists(instruction *Instruction) bool {
	exists := false
	for _, currInstruction := range *b {
		if exists = currInstruction.Equals(instruction); exists {
			break
		}
	}
	return exists
}
