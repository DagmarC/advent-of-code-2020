package datafile

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/DagmarC/advent-of-code-2020/utils"
)

func LoadAndParseTask3() [][]string {

	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// Find the length of the lines.
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		scanner.Text()
		lineNumber++
	}
	// Reset File pointer back to the beginning.
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil
	}

	// InitializeDistribution slice of slices and stores the file into the slice of slices.
	forestMap := make([][]string, lineNumber)
	scanner = bufio.NewScanner(file)
	lineNumber = 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, l := range line {
			forestMap[lineNumber] = append(forestMap[lineNumber], string(l))
		}
		lineNumber++
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
	return forestMap
}
