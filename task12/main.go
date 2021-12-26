package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/task12/ship"
)

// At the end of these instructions, the ship's Manhattan distance (sum of the absolute values of its
//east/west position and its north/south position) from its starting position is 17 + 8 = 25.

func main() {

	instructions, err := datafile.LoadFileTask12()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instructions)

	manhattanDistance, err := ship.CalculateManhattanDistance(instructions, false)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("T1: Manhattan distance of the ship after taking instructions: ", manhattanDistance)

	manhattanRelativeDistance, err := ship.CalculateManhattanDistance(instructions, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("T2: Manhattan relative distance of the ship from waypoint after taking instructions: ",
		manhattanRelativeDistance)

}
