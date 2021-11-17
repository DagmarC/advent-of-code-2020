// https://adventofcode.com/2020/day/13#part2
package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task13/busStation"
	"github.com/DagmarC/codeOfAdvent/task13/chineseReminderTheory"
	"github.com/DagmarC/codeOfAdvent/task13/task2"
)

func main() {

	chineseReminderTheory.Test()
	fmt.Println("+++++++++++++++++++++++++++++++++++++")

	allBusses, leavingTime, err := datafile.LoadFileTask13()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allBusses, leavingTime)
	earliestBus := findNearestBusDeparture(leavingTime, allBusses)

	fmt.Printf("What is the ID of the earliest bus you can take to the airport multiplied by the number of " +
		"minutes you'll need to wait for that bus?\n")
	fmt.Printf("The bus is: %v and the result is %v.\n", earliestBus, int(earliestBus.BusId)*earliestBus.Offset)

	fmt.Println("+++++++++++++++++++++++++++++++++++++")

	// _________________________________________________________________________________________________________________
	// _________________________________________________________________________________________________________________

	earliestTimestampWithOffset := task2.FindNearestTimestamp(&allBusses)
	fmt.Println("What is the earliest timestamp such that all of the listed bus IDs depart at offsets matching their " +
		"positions in the list?")
	fmt.Println("Task2 result:", earliestTimestampWithOffset)
}

func findNearestBusDeparture(earliestDepartureTime busStation.DepartureTime, busses []busStation.NearestBusDeparture) busStation.NearestBusDeparture {

	var earliestBus busStation.NearestBusDeparture
	earliestBus.Offset = 1000000000

	for _, bus := range busses {
		nextDeparture, waitTime := bus.BusId.CalculateNextDepartureFrom(earliestDepartureTime)
		if waitTime < earliestBus.Offset {
			earliestBus.Update(bus.BusId, nextDeparture, waitTime)
		}
	}

	return earliestBus
}
