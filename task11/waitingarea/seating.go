package waitingarea

import "fmt"

// Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable
//and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a
//given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat).

//The following rules are applied to every seat simultaneously:
// ------------------------------------------------------------
//--If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
//--If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
//Otherwise, the seat's state does not change.
//--Floor (.) never changes; seats don't move, and nobody sits on the floor.

func StabilizeSeats(layout *SeatLayout) SeatLayout {

	for {
		// Copy is needed, because the rules applies on the original array whole time, at the end it will be reassigned.
		newLayout := createLayoutShallowCopy(layout)
		changeCounter := 0 // Reset the changeCounter.

		for rowIndex, _ := range *layout {
			for lineIndex, _ := range (*layout)[rowIndex] {
				currentSeat := (*layout)[rowIndex][lineIndex]

				newSeat := getSeatOccupation(currentSeat, layout, rowIndex, lineIndex) // Apply rules on the original layout.

				if newSeat != currentSeat {
					newLayout[rowIndex][lineIndex] = newSeat // Change newLayout.
					changeCounter++
				}
			}
		}
		if changeCounter == 0 {
			fmt.Println("Breaking")
			break // No more seat adjustments are needed.
		}
		layout = &newLayout // Now reassign the layout to newLayout.
	}
	return *layout
}

func createLayoutShallowCopy(layout *SeatLayout) SeatLayout {
	newLayout := make(SeatLayout, len(*layout))
	for i := range *layout {
		newLayout[i] = make([]SeatType, len((*layout)[i]))
		copy(newLayout[i], (*layout)[i]) // To use copy() properly, len() must be the same.
	}
	return newLayout
}

func getSeatOccupation(currentSeat SeatType, layout *SeatLayout, row, line int) SeatType {
	if currentSeat == Floor {
		return Floor // No change is needed here.
	}

	nAdjacentOccupiedSeats := getNumberAdjacentOccupiedSeats(layout, row, line)

	if nAdjacentOccupiedSeats == 0 && currentSeat == EmptySeat {
		return OccupiedSeat
	}
	if nAdjacentOccupiedSeats >= 4 && currentSeat == OccupiedSeat {
		return EmptySeat
	}
	return currentSeat
}

func getNumberAdjacentOccupiedSeats(layout *SeatLayout, row int, line int) int {
	occupied := 0
rowLoop:
	for r := row - 1; r <= row+1; r++ {
		for l := line - 1; l <= line+1; l++ {

			if r == row && l == line {
				continue // Do not count the current seat position, only its adjacent seats.
			}
			if r < 0 || r >= len(*layout) {
				continue rowLoop // r index out of range.
			}
			if l < 0 || l >= len((*layout)[row]) {
				continue // l index out of range.
			}
			if (*layout)[r][l] == OccupiedSeat {
				occupied++
			}
		}
	}
	return occupied
}
