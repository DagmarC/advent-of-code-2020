package main

import (
	"fmt"
	"testing"

	"github.com/DagmarC/codeOfAdvent/task9/dataqueue"
)

type testData struct {
	result map[int64]int
	input  []int
}

var testInputs = []testData{
	{
		map[int64]int{2020: 436, 30000000: 175594},
		[]int{0, 3, 6},
	},
	{
		result: map[int64]int{2020: 1, 30000000: 2578},
		input:  []int{1, 3, 2},
	},
	{
		map[int64]int{2020: 1836, 30000000: 362},
		[]int{3, 1, 2},
	},
}

func TestStartGame(t *testing.T) {
	for _, data := range testInputs {
		testElfGame := createElfGame(data.input)

		for stopGame, expectedResult := range data.result {

			testElfGame.saidNumbers = make(map[int]*dataqueue.Queue, 0) // Reset the game queue.

			result := testElfGame.Start(stopGame)
			if result != expectedResult {
				t.Errorf("Expected result %d, got %d.", expectedResult, result)
				t.Fail()
			}
			fmt.Printf("Stop game after %d for input %v - the Number is %d\n", stopGame, data.input, result)
		}
	}
}
