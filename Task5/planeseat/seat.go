package planeseat

import (
	"math"
)

type Seat struct {
	row    string
	column string
	seatID int
}

// SetRowIDString sets row.
func (s *Seat) SetRowIDString(id string) {
	s.row = id
}

// SetColumnIDString sets column.
func (s *Seat) SetColumnIDString(id string) {
	s.column = id
}

// SetSeatID Every seat also has a unique seat ID: multiply the row by 8, then add the column.
// In this example, the seat has ID 44 * 8 + 5 = 357.
func (s *Seat) SetSeatID() int {
	s.seatID = 8*s.GetSeatRowValue() + s.GetSeatColumnValue()
	return s.seatID
}

func (s *Seat) GetSeatID() int {
	return s.seatID
}

// GetSeatRowValue Will obtain value via converting from binary into decimal representation:
//// FBFBBFF - 7 digit string, where F represents 0 and B represents 1.
func (s *Seat) GetSeatRowValue() int {
	power := 0.0
	result := 0.0

	for i := len(s.row) - 1; i >= 0; i-- {
		switch string(s.row[i]) {
		case "B":
			result += math.Pow(2, power)
		}
		power++
	}
	return int(result)
}

// GetSeatColumnValue Will obtain value via converting from binary into decimal representation:
// RLR - 3 digit string, where L represents 0 and R represents 1.
func (s *Seat) GetSeatColumnValue() int {
	power := 0.0
	result := 0.0

	for i := len(s.column) - 1; i >= 0; i-- {
		switch string(s.column[i]) {
		case "R":
			result += math.Pow(2, power)
		}
		power++
	}
	return int(result)
}
