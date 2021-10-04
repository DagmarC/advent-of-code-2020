package main

import (
	"errors"
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task9/dataqueue"
	"log"
)

func main() {
	numbers, err := datafile.LoadFileTask9()
	if err != nil {
		log.Fatal(err)
	}
	offset := 25
	//preambleTest := createPreamble(5, &numbers)
	preamble := createPreamble(offset, &numbers)

	resultNumber, index, err := findFirstWrongNumber(preamble, &numbers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TASK 1 result: %d at index %d.\n", resultNumber, index)

	sumNumbers, err := findConsecutiveNumbersForSum(resultNumber, index, &numbers)
	fmt.Printf("The consecutive elements %v that sum up to %d and the sum of max and min is %d.\n", sumNumbers,
		resultNumber, sumNumbers.Max()+sumNumbers.Min())
}

func createPreamble(size int, numbers *[]int) *dataqueue.Queue {
	preamble := dataqueue.Queue{}
	for i := 0; i < size; i++ {
		preamble.Enqueue((*numbers)[i])
	}
	return &preamble
}

// TASK 1
//In this example, after the 5-number preamble, almost every number is the sum of two of the previous 5 numbers;
//the only number that does not follow this rule is 127.
//The first step of attacking the weakness in the XMAS data is to find the first number in the list (after the preamble)
//which is not the sum of two of the 25 numbers before it. What is the first number that does not have this property?

// findFirstWrongNumber returns the first number that will not match the sum of the any 2 distinct previous numbers.
func findFirstWrongNumber(preamble *dataqueue.Queue, numbers *[]int) (int, int, error) {

	offset := preamble.Length()
	fmt.Println(offset)
	for i := offset; i < len(*numbers); i++ {
		currentNumber := (*numbers)[i]
		if match := findSumForNumber(preamble, currentNumber); !match {
			return currentNumber, i, nil
		}
		// Remove the 1st element - no longer actual.
		_, err := preamble.Dequeue()
		if err != nil {
			return 0, -1, err
		}
		// Add the current element to have updated preamble.
		preamble.Enqueue(currentNumber)
	}
	return 0, -1, errors.New("no wrong number was detected")
}

// TASK 2
// findConsecutiveNumbersForSum find a contiguous set of at least two numbers in your list which sum to the
// invalid number from TASK 1.
func findConsecutiveNumbersForSum(findSum int, maxIndex int, numbers *[]int) (dataqueue.Queue, error) {
	tempSum := 0
	sumQueue := dataqueue.Queue{}

	for i := 0; i < maxIndex; i++ {
		n := (*numbers)[i]
		tempSum += n
		sumQueue.Enqueue((*numbers)[i])

		for tempSum > findSum {
			toRemove, err := sumQueue.Dequeue()
			if err != nil {
				return dataqueue.Queue{}, err
			}
			tempSum -= toRemove
		}
		// Find the consecutive numbers that findSum up to the number-findSum.
		if tempSum == findSum {
			return sumQueue, nil
		}
	}
	return dataqueue.Queue{}, errors.New("no consecutive numbers will Sum up to the wanted number")
}

func findSumForNumber(preamble *dataqueue.Queue, number int) bool {
	for i := 0; i < preamble.Length(); i++ {
		for j := 0; j < preamble.Length(); j++ {
			if i == j {
				continue
			}
			x, y := preamble.Get(i), preamble.Get(j)

			if x+y == number {
				return true
			}
		}
	}
	return false
}
