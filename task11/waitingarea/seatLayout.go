package waitingarea

import "fmt"

type SeatType byte

const (
	Floor        SeatType = '.'
	EmptySeat    SeatType = 'L'
	OccupiedSeat SeatType = '#'
)

type SeatLayout [][]SeatType

func (s *SeatLayout) Initialize(len int) {
	*s = make(SeatLayout, len)
}

func (s *SeatLayout) InitializeLine(row, capacity int) {
	(*s)[row] = make([]SeatType, 0, capacity)
}

func (s *SeatLayout) Print() {
	for _, row := range *s {
		for _, s := range row {
			fmt.Print(string(s))
		}
		fmt.Println()
	}
}

func (s *SeatLayout) OccupiedSeats() int {
	count := 0
	for _, row := range *s {
		for _, s := range row {
			if s == OccupiedSeat {
				count++
			}
		}
	}
	return count
}
