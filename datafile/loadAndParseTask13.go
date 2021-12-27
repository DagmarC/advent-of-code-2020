package datafile

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task13/busStation"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

func LoadFileTask13() ([]busStation.NearestBusDeparture, busStation.DepartureTime, error) {

	file, err := os.Open(path.Join(utils.GetWd(), constants.Input))
	if err != nil {
		return nil, -1, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	allBusses := make([]busStation.NearestBusDeparture, 0)
	var leavingTime busStation.DepartureTime

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, -1, err
		}
		// First row - the earliest timestamp you could depart
		if row == 0 {
			time, err := strconv.Atoi(line)
			if err != nil {
				return nil, -1, err
			}
			leavingTime = busStation.DepartureTime(time)
		} else {
			err = parseLineAndAppendBusses(line, &allBusses)
			if err != nil {
				return nil, -1, err
			}
		}
		row++
	}
	if scanner.Err() != nil {
		return nil, -1, err
	}
	return allBusses, leavingTime, nil
}

func parseLineAndAppendBusses(line string, allBusses *[]busStation.NearestBusDeparture) error {
	busses := strings.Split(line, ",")
	for offset, bus := range busses {
		busId, err := strconv.Atoi(bus)
		if err != nil {
			continue //x
		}
		*allBusses = append(*allBusses, busStation.NearestBusDeparture{
			BusId:  busStation.BusId(busId),
			Offset: offset,
		})
	}
	return nil
}
