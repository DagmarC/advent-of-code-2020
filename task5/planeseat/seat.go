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
	s.seatID = 8*s.SeatRowValue() + s.SeatColumnValue()
	return s.seatID
}

func (s *Seat) SeatID() int {
	return s.seatID
}

// SeatRowValue Will obtain value via converting from binary into decimal representation:
//// FBFBBFF - 7 digit string, where F represents 0 and B represents 1.
func (s *Seat) SeatRowValue() int {
	power := 0.0
	result := 0.0

	for i := len(s.row) - 1; i >= 0; i-- {
		if string(s.row[i]) == "B" {
			result += math.Pow(2, power)
		}
		power++
	}
	return int(result)
}

// SeatColumnValue Will obtain value via converting from binary into decimal representation:
// RLR - 3 digit string, where L represents 0 and R represents 1.
func (s *Seat) SeatColumnValue() int {
	power := 0.0
	result := 0.0

	for i := len(s.column) - 1; i >= 0; i-- {
		if string(s.column[i]) == "R" {
			result += math.Pow(2, power)
		}
		power++
	}
	return int(result)
}
