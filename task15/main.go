package main

import (
	"fmt"
	"path"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

func main() {
	input := datafile.ReadLineOfNumbers(path.Join(utils.GetWd(), constants.Input))

	fmt.Println("-------TASK 1-------")
	game := createElfGame(input)
	endGameNumber := game.Start(2020)
	fmt.Println("Given your starting numbers, what will be the 2020th number spoken?", endGameNumber)

	// Reset the game queue.
	game.saidNumbers = make(map[int]*utils.Queue, 0)

	fmt.Println("-------TASK 2-------")
	endGameNumber = game.Start(30000000)
	fmt.Println("Given your starting numbers, what will be the 2020th number spoken?", endGameNumber)
}

// ElfGame saidNumbers: key:value a pair of number:position
type ElfGame struct {
	input       []int
	saidNumbers map[int]*utils.Queue
	lastNumber  int
}

func createElfGame(input []int) *ElfGame {
	return &ElfGame{
		input:       input,
		saidNumbers: make(map[int]*utils.Queue, 0),
	}
}

func (e *ElfGame) Start(stopGame int) int {
	currentNumber := 0

	for pos := 0; pos < stopGame; pos++ {
		if pos < len(e.input) {
			currentNumber = e.input[pos]
		} else {
			currentNumber = getNextNumber(e)
		}
		addLastSaidNumber(e, currentNumber, pos+1) // Save Current number as the last said number.
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

func addLastSaidNumber(e *ElfGame, number int, position int) {
	e.lastNumber = number

	if _, ok := e.saidNumbers[number]; !ok {
		e.saidNumbers[number] = &utils.Queue{} // Initialize queue, if the number is not present in map.
	}
	e.saidNumbers[number].Enqueue(position)
}
