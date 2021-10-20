package ship

import (
	"fmt"
	"math"
)

// CalculateManhattanDistance relative: false is TASK1, relative: true is TASK2
func CalculateManhattanDistance(instructions []Instruction, relative bool) (float64, error) {
	ship := CreateShip()

	for _, instruction := range instructions {
		var err error
		if relative {
			err = ship.moveShipRelatively(&instruction)
		} else {
			err = ship.performInstruction(&instruction)
		}
		if err != nil {
			return 0, err
		}
	}
	fmt.Println("SHIP", ship)
	return math.Abs(ship.coordinates[0]) + math.Abs(ship.coordinates[1]), nil
}
