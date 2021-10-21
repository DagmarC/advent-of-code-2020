package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task13/busStation"
	"log"
)

func main() {

	allBusses, leavingTime, err := datafile.LoadFileTask13()
	if err != nil {
		log.Fatal(err)
	}
	earliestBus := findNearestBusDeparture(leavingTime, allBusses)
	fmt.Printf("What is the ID of the earliest bus you can take to the airport multiplied by the number of " +
		"minutes you'll need to wait for that bus?\n")
	fmt.Printf("The bus is: %v and the result is %v.\n", earliestBus, int(earliestBus.BusId)*earliestBus.WaitTime)
}

func findNearestBusDeparture(earliestDepartureTime busStation.DepartureTime, busses []busStation.Bus) busStation.NearestBusDeparture {

	var earliestBus busStation.NearestBusDeparture
	earliestBus.WaitTime = 100000000000000000

	for _, busId := range busses {
		nextDeparture, waitTime := busId.CalculateNextDepartureFrom(earliestDepartureTime)
		if waitTime < earliestBus.WaitTime {
			earliestBus.Update(busId, nextDeparture, waitTime)
		}
	}

	return earliestBus
}
