package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task11/waitingarea"
	"log"
)

func main() {
	// TASK 1
	seatLayout, err := datafile.LoadFileTask11()
	if err != nil {
		log.Fatal(err)
	}
	seatLayout = waitingarea.StabilizeSeats(&seatLayout, 1)
	fmt.Println("TASK 1: Number of occupied seats:", seatLayout.OccupiedSeats())

	// TASK 2
	seatLayout, err = datafile.LoadFileTask11()
	if err != nil {
		log.Fatal(err)
	}
	seatLayout = waitingarea.StabilizeSeats(&seatLayout, 2)
	fmt.Println("TASK 2: Number of occupied seats:", seatLayout.OccupiedSeats())
}
