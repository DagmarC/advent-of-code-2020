package datafile

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/DagmarC/advent-of-code-2020/task5/planeseat"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

// LoadFileTask5 loads, parses creates and returns *[]planeseat.Seat
func LoadFileTask5() (*[]*planeseat.Seat, error) {

	file, err := os.Open(utils.GetInputPath())
  	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	planeSeats := make([]*planeseat.Seat, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Save parsed line into the struct.
		seat, err := parseLineGetSeat(line)
		if err != nil {
			return nil, err
		}
		planeSeats = append(planeSeats, seat)
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return &planeSeats, nil
}

func parseLineGetSeat(line string) (*planeseat.Seat, error) {
	seat := planeseat.Seat{}
	if len(line) != 10 {
		return &seat, errors.New("wrong number format. Must have 10 charachters")
	}
	seat.SetRowIDString(line[:7])
	seat.SetColumnIDString(line[7:])

	return &seat, nil
}
