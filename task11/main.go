package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task11/waitingarea"
	"log"
)

func main() {
	seatLayout, err := datafile.LoadFileTask11()
	if err != nil {
		log.Fatal(err)
	}
	seatLayout = waitingarea.StabilizeSeats(&seatLayout)
	seatLayout.Print()
	fmt.Println("TASK 1: Number of occupied seats:", seatLayout.OccupiedSeats())

}
