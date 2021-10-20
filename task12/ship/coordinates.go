package ship

import (
	"errors"
	"fmt"
	"math"
)

const (
	N Direction = "N"
	E Direction = "E" // Default one
	S Direction = "S"
	W Direction = "W"
	F Direction = "F"
	R Direction = "R"
	L Direction = "L"
)

// Direction is the Direction in which you should move -> N(north), F(Forward), ...
type Direction string

// Coordinates Cartesian coordinates [x, y]
type Coordinates [2]float64

type Compass map[Direction]float64

var compass = make(Compass, 4)

// The waypoint starts 10 units east and 1 unit north relative to the ship.
//The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.
var waypoint Coordinates = [2]float64{10, 1}

func (c *Compass) Initialize() {
	compass[N] = 0
	compass[E] = 90
	compass[S] = 180
	compass[W] = 270
}

func (c *Coordinates) Directions() []Direction {
	directions := make([]Direction, len(c))

	if c[0] >= 0 {
		directions[0] = E
	} else {
		directions[0] = W
	}
	if c[1] >= 0 {
		directions[1] = N
	} else {
		directions[1] = S
	}

	return directions
}

// UpdateDirection will update the coordinates respectively via directions.
// waypoint stands for axis [+-x, +-y] positions (E/W, N/S).
// So if the waypoint is [10, 4] meaning (N 10, E 4) and directions [S, E] then new waypoint will be [4,-10] (E 4, S 10),
// meaning that the East position 10 moved to south ==> change sign and place in waypoint [_, -10] and
// The North position 4 moved E so the sign remains the same and the position in the waypoint changes ==> [4, _].
// Result [4, -10] stands for East 4, South 10.
func (c *Coordinates) UpdateDirection(directions ...Direction) {
	currentDirections := c.Directions()
	tmpWaypoints := Coordinates{}

	for i, direction := range directions {
		if currentDirections[i] == direction {
			continue
		}
		switch direction {
		case S:
			tmpWaypoints[1] = -math.Abs(c[i])
		case N:
			tmpWaypoints[1] = math.Abs(c[i])
		case W:
			tmpWaypoints[0] = -math.Abs(c[i])
		case E:
			tmpWaypoints[0] = math.Abs(c[i])
		}
	}
	*c = tmpWaypoints
}

func (c *Coordinates) Update(c1, c2 float64) {
	c[0] += c1
	c[1] += c2
}

func calculateNewDirections(degrees float64, directions ...Direction) ([]Direction, error) {
	newDirections := make([]Direction, 0)

	for _, direction := range directions {
		newDegrees := math.Remainder(compass[direction]+degrees, 360.0)
		newDirection, err := compass.GetByValue(newDegrees)
		if err != nil {
			return newDirections, err
		}
		newDirections = append(newDirections, newDirection)
	}
	return newDirections, nil
}

func (c *Compass) GetByValue(degrees float64) (Direction, error) {

	if degrees < 0 {
		degrees = covertNegativeDegrees(degrees)
	}
	for key, value := range compass {
		if value == degrees {
			return key, nil
		}
	}
	fmt.Println(degrees)
	return N, errors.New("wrong degrees")
}

func covertNegativeDegrees(degrees float64) float64 {
	// Opposite directions:
	if degrees == -90 {
		degrees = 270
	}
	if degrees == -270 {
		degrees = 90
	}
	if degrees < 0 {
		degrees *= -1
	}
	return degrees
}

// Instruction is the one step in ship navigation.
type Instruction struct {
	action Direction
	value  float64
}

func CreateInstruction(cardinalDirection string, value float64) *Instruction {
	return &Instruction{action: Direction(cardinalDirection), value: value}
}

// adjustShipCoordinates Axis x, y --> North: +y, South: -y, West: -x, East: +x
// TASK 1
func (i *Instruction) adjustShipCoordinates(ship *Ship) error {
	switch i.action {
	case N:
		ship.coordinates[0] += i.value
	case S:
		ship.coordinates[0] -= i.value
	case E:
		ship.coordinates[1] += i.value
	case W:
		ship.coordinates[1] -= i.value
	case F:
		i.action = ship.cardinalDirection
		err := i.adjustShipCoordinates(ship)
		if err != nil {
			return err
		}
	default:
		return errors.New("wrong instruction")
	}
	return nil
}
