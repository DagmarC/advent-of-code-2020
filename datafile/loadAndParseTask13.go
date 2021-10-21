package datafile

import (
	"bufio"
	"fmt"
	"github.com/DagmarC/codeOfAdvent/constants"
	"github.com/DagmarC/codeOfAdvent/task13/busStation"
	"log"
	"os"
	"regexp"
	"strconv"
)

func LoadFileTask13() ([]busStation.Bus, busStation.DepartureTime, error) {

	file, err := os.Open(constants.Task13)
	if err != nil {
		return nil, -1, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	allBusses := make([]busStation.Bus, 0)
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
	fmt.Println(allBusses, leavingTime)
	return allBusses, leavingTime, nil
}

func parseLineAndAppendBusses(line string, allBusses *[]busStation.Bus) error {
	// line 19,x,x,21,...
	reg := regexp.MustCompile("([0-9]+)[,x]*")
	// Find numbers: matches [[19,x,x,x,x,x,x,x,x, 19], [..., N], ...}
	matches := reg.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		if len(match) == 2 {
			busId, err := strconv.Atoi(match[1])
			if err != nil {
				return err
			}
			*allBusses = append(*allBusses, busStation.Bus(busId))
		}
	}
	return nil
}
