package datafile

import (
	"bufio"
	"errors"

	"log"
	"os"
	"strconv"

	"github.com/DagmarC/advent-of-code-2020/task10/adapterpkg"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

func LoadFileTask10() ([]adapterpkg.Adapter, error) {

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

	allAdapters := make([]adapterpkg.Adapter, 0)
	// Treat the charging outlet near your seat as having an effective joltage rating of 0.
	allAdapters = append(allAdapters, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}

		err = parseLineAndAppendAdapters(line, &allAdapters)
		if err != nil {
			return nil, err
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return allAdapters, nil
}

func parseLineAndAppendAdapters(line string, adapters *[]adapterpkg.Adapter) error {
	n, err := strconv.Atoi(line)
	if err != nil {
		return errors.New("NaN")
	}
	var adapter = adapterpkg.Adapter(n)
	*adapters = append(*adapters, adapter)
	return nil
}
