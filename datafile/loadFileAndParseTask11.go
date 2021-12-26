package datafile

import (
	"bufio"
	"fmt"

	"log"
	"os"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task11/waitingarea"
)

func LoadFileTask11() (waitingarea.SeatLayout, error) {

	file, err := os.Open(constants.Task11)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	seatLayout := waitingarea.SeatLayout{}
	seatLayout.Initialize(0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}

		seatLayout = append(seatLayout, []waitingarea.SeatType{})
		seatLayout.InitializeLine(row, len(line))

		err = parseSeatLayout(line, &seatLayout[row])
		if err != nil {
			return nil, err
		}
		row++
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return seatLayout, nil
}

func parseSeatLayout(line string, seatLayoutRow *[]waitingarea.SeatType) error {
	for _, character := range line {
		seatType := waitingarea.SeatType(character)
		switch seatType {
		case waitingarea.Floor:
			*seatLayoutRow = append(*seatLayoutRow, waitingarea.Floor)
		case waitingarea.OccupiedSeat:
			*seatLayoutRow = append(*seatLayoutRow, waitingarea.OccupiedSeat)
		case waitingarea.EmptySeat:
			*seatLayoutRow = append(*seatLayoutRow, waitingarea.EmptySeat)
		default:
			return fmt.Errorf("this seat %v is not a valiid option", seatType)
		}
	}
	return nil
}
