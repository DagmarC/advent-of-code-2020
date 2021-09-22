// For example, consider just the first seven characters of FBFBBFFRLR:
//FBFBBFFRLR
//0123456789
//
//Start by considering the whole range, rows 0 through 127.
//F means to take the lower half, keeping rows 0 through 63.
//B means to take the upper half, keeping rows 32 through 63.
//F means to take the lower half, keeping rows 32 through 47.
//B means to take the upper half, keeping rows 40 through 47.
//B keeps rows 44 through 47.
//F keeps rows 44 through 45.
//The final F keeps the lower of the two, row 44.
//
//Start by considering the whole range, columns 0 through 7.
//R means to take the upper half, keeping columns 4 through 7.
//L means to take the lower half, keeping columns 4 through 5.
//The final R keeps the upper of the two, column 5.

package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/Task5/planeseat"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"math"
	"sort"
)

// As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?
func main() {
	seats, err := datafile.LoadFileTask5()
	if err != nil {
		fmt.Println(err)
	}
	maxSeatID := int(math.Inf(-1))
	for _, s := range *seats {
		seatID := s.SetSeatID()
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Println("Max Seat ID is:", maxSeatID)

	// PART II
	mySeat := getMissingSeatID(*seats)
	fmt.Println("My lost seat number is:", mySeat)

}

// PART II
func getMissingSeatID(seats []*planeseat.Seat) int {
	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(seats, func(i, j int) bool {
		return seats[i].GetSeatID() < seats[j].GetSeatID()
	})
	currentSeat := seats[0].GetSeatID()
	for _, s := range seats {
		if currentSeat != s.GetSeatID() {
			return currentSeat
		}
		currentSeat++
	}
	return -1
}
