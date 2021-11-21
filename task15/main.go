package main

import (
	"fmt"

	"github.com/DagmarC/codeOfAdvent/constants"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task9/dataqueue"
	"github.com/DagmarC/codeOfAdvent/utils"
)

func main() {
	input := datafile.ReadLineOfNumbers(constants.Task15)

	fmt.Println("-------TASK 1-------")
	game := createElfGame(input)
	endGameNumber := game.Start(2020)
	fmt.Println("Given your starting numbers, what will be the 2020th number spoken?", endGameNumber)

	// Reset the game queue.
	game.saidNumbers = make(map[int]*dataqueue.Queue, 0)

	fmt.Println("-------TASK 2-------")
	endGameNumber = game.Start(30000000)
	fmt.Println("Given your starting numbers, what will be the 2020th number spoken?", endGameNumber)
}

// ElfGame saidNumbers: key:value a pair of number:position
type ElfGame struct {
	input       []int
	saidNumbers map[int]*dataqueue.Queue
	lastNumber  int
}

func createElfGame(input []int) *ElfGame {
	return &ElfGame{
		input:       input,
		saidNumbers: make(map[int]*dataqueue.Queue, 0),
	}
}

func (e *ElfGame) Start(stopGame int64) int {
	currentNumber := 0

	var pos = int64(1)
	for _, n := range e.input {
		addLastSaidNumber(e, n, pos)
		pos++
	}
	for pos <= stopGame {
		currentNumber = getNextNumber(e)
		addLastSaidNumber(e, currentNumber, pos) // Save Current number as the last said number.
		pos++
	}
	return currentNumber
}

// getNextNumber gets the number to be spoken according to the last said number, if the number was spoken only once - return 0
func getNextNumber(e *ElfGame) int {
	queue := e.saidNumbers[e.lastNumber]
	if queue.Length() == 1 {
		return 0
	}
	firstNumber, err := queue.Dequeue()
	utils.Check(err)

	return queue.Get(0) - firstNumber
}

func addLastSaidNumber(e *ElfGame, number int, position int64) {
	e.lastNumber = number

	if _, ok := e.saidNumbers[number]; !ok {
		e.saidNumbers[number] = &dataqueue.Queue{} // Initialize queue, if the number is not present in map.
	}
	e.saidNumbers[number].Enqueue(int(position))

}
