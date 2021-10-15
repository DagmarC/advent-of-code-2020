package waitingarea

import "fmt"

func StabilizeSeats(layout *SeatLayout, task int) SeatLayout {

	for {
		// Copy is needed, because the rules applies on the original array whole time, at the end it will be reassigned.
		newLayout := createLayoutShallowCopy(layout)
		changeCounter := 0 // Reset the changeCounter.
		var newSeat SeatType

		for rowIndex := range *layout {
			for lineIndex := range (*layout)[rowIndex] {
				currentSeat := (*layout)[rowIndex][lineIndex]

				if task == 1 {
					newSeat = getSeatOccupation(currentSeat, layout, rowIndex, lineIndex, 4, task) // Apply rules on the original layout.
				} else {
					newSeat = getSeatOccupation(currentSeat, layout, rowIndex, lineIndex, 5, task) // Apply rules on the original layout.
				}

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

func getSeatOccupation(currentSeat SeatType, layout *SeatLayout, row int, line int, ruleToEmpty int, task int) SeatType {
	if currentSeat == Floor {
		return Floor // No change is needed here.
	}
	var nOccupiedSeats = 0

	if task == 1 {
		// In all 4 directions.
		nOccupiedSeats = getNumberAdjacentOccupiedSeats(layout, row, line)
	} else {
		allDirs := Directions{}
		allDirs.Initialize()
		for _, dir := range allDirs.allDirections {
			nOccupiedSeats += getOccupiedSeatsT2(layout, row, line, dir)
		}
	}

	if nOccupiedSeats == 0 && currentSeat == EmptySeat {
		return OccupiedSeat
	}
	if nOccupiedSeats >= ruleToEmpty && currentSeat == OccupiedSeat {
		return EmptySeat
	}
	return currentSeat
}
