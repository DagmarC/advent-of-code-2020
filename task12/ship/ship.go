package ship

// Ship has cardinalDirection - what cardinalDirection it is facing to and current coordinates.
type Ship struct {
	cardinalDirection Direction
	coordinates       Coordinates
}

// TASK 1
func (s *Ship) turnShip(degrees float64, side Direction) error {

	if side == L {
		degrees = -degrees
	}
	newDirection, err := calculateNewDirections(degrees, s.cardinalDirection)
	if err != nil {
		return err
	}
	s.cardinalDirection = newDirection[0]
	return nil
}

// performInstruction TASK 1
func (s *Ship) performInstruction(instruction *Instruction) error {

	if instruction.action == R || instruction.action == L {
		err := s.turnShip(instruction.value, instruction.action)
		if err != nil {
			return err
		}
		return nil
	}

	err := instruction.adjustShipCoordinates(s)
	if err != nil {
		return err
	}

	return nil
}

// moveShipRelatively TASK 2
func (s *Ship) moveShipRelatively(instruction *Instruction) error {

	if instruction.action == F {
		s.moveForward(instruction)
		return nil
	}

	switch instruction.action {
	case R:
		newDirections, err := calculateNewDirections(instruction.value, waypoint.Directions()...)
		if err != nil {
			return err
		}
		waypoint.UpdateDirection(newDirections...)
	case L:
		newDirections, err := calculateNewDirections(-instruction.value, waypoint.Directions()...)
		if err != nil {
			return err
		}
		waypoint.UpdateDirection(newDirections...)
	case N:
		waypoint.Update(0, instruction.value)
	case S:
		waypoint.Update(0, -instruction.value)
	case W:
		waypoint.Update(-instruction.value, 0)
	case E:
		waypoint.Update(instruction.value, 0)

	}
	return nil
}

func (s *Ship) moveForward(instruction *Instruction) {
	s.coordinates[0] += instruction.value * waypoint[0]
	s.coordinates[1] += instruction.value * waypoint[1]
}

func CreateShip() *Ship {
	compass.Initialize()
	return &Ship{
		cardinalDirection: E,
		coordinates:       Coordinates{},
	}
}
