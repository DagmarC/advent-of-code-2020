package waitingarea

// **********************************TASK 2********************************************
// Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an
//occupied seat to become empty (rather than four or more from the previous rules). The other rules still apply:
//empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.

func getOccupiedSeatsT2(layout *SeatLayout, row, line int, dir direction) int {
	count := 0

	rowDirection := dir[0]
	lineDirection := dir[1]

	for r, l := row+rowDirection, line+lineDirection; ; r, l = r+rowDirection, l+lineDirection {
		// OUT OF INDEX CONDITION
		if r < 0 || r >= len(*layout) || l < 0 || l >= len((*layout)[row]) {
			break
		}

		if (*layout)[r][l] == EmptySeat {
			break // No other seats are seen behind the empty seat.
		}
		if (*layout)[r][l] == OccupiedSeat {
			count++
			break
		}
	}
	return count
}
