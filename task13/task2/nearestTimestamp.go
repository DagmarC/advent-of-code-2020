package task2

import (
	"fmt"

	"github.com/DagmarC/codeOfAdvent/task13/busStation"
)

// FindNearestTimestamp the earliest timestamp such that all the listed bus IDs depart at offsets matching
// their positions in the list?
// Will go through the list of busses and each time it will find the common timestamp intersection of two busses that
// will then be used with the next bus.
// Logic behind:
// busId with offset 0 (0th position) must be divisible by timestamp t. All the next busses with specified
// offsets are congruent -> t is congruent to busBId - (*busses)[i].Offset % busBId (mod busBId), which in fact means
// that after t is divided by bus[i] the reminder is calculated as busBId - its offset (mod busId, if offset > busId)
// It wil one by one combine busses and use the common timestamp in the next iteration.
func FindNearestTimestamp(busses *[]busStation.NearestBusDeparture) int {

	busATimestamp := int((*busses)[0].BusId)
	difference := busATimestamp
	lastElement := false

	for i := 1; i < len(*busses); i++ {

		busBId := int((*busses)[i].BusId)
		reminder := busBId - (*busses)[i].Offset%busBId

		if i == len(*busses)-1 {
			lastElement = true
		}
		busATimestamp, difference = findCommonTimestamp(busATimestamp, busBId, reminder, difference, lastElement)
	}

	return busATimestamp
}

// findCommonTimestamp This method does 2 things.
// Firstly it is looking for a busA timestamp such that busA % busB == reminder, and it does it exactly 2 times
// because it is looking for a difference in between two values that matches the condition.
// When it matches two different busA timestamps and calculates its difference, the result is the value of busA timestamp
// after two matches as well as the difference, which will in the next iteration  be used as increment to be more effective.
// If it is the last element, we can skip the step for looking for the difference, because we already reached the destination
// and the final result - the earliest timestamp is already found after the first good catch of busA % busB == reminder.
func findCommonTimestamp(busA, busB, reminder, increment int, lastElement bool) (int, int) {

	fmt.Printf("a=%v, b=%v, r=%v, inc:=%v\n", busA, busB, reminder, increment)

	var differenceOfTimestamps int
	// to catch the difference, you need to have the stop value to reach at least 2
	stop := 0
	oldBusValue := 0

	for {
		if busA%busB == reminder {

			differenceOfTimestamps = busA - oldBusValue
			stop++ // Need at least two matches.

			if stop >= 2 || lastElement {
				break // if lastElement then the result is found.
			}
			oldBusValue = busA
		}
		busA += increment

	}
	return busA, differenceOfTimestamps
}
