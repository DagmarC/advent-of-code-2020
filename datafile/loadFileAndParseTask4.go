package datafile

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/task4/passp"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

const (
	space     = " "
	semicolon = ":"
)

func LoadAndParseTask4() []*passp.Passport {

	var passports []*passp.Passport

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

	scanner := bufio.NewScanner(file)
	passport := passp.CreatePassport()
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		// If blank line -> Passport loading ends, append the passport and be ready for the new one.
		if line == "" {
			passports = append(passports, passport)
			passport = passp.CreatePassport()
			continue
		}
		err := parsePassport(passport, line)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Append the last password - if and only if the last line is not the blank line.
	if line != "" {
		passports = append(passports, passport)
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return passports
}

// parsePassport will parse line: hgt:154cm eyr:2030 and store it into the passport variable.
func parsePassport(passport *passp.Passport, line string) error {
	// [hgt:154cm, eyr:2030]
	passData := strings.Split(line, space)

	for _, data := range passData {
		// [hgt, 154cm]
		passFields := strings.Split(data, semicolon)
		if len(passFields) != 2 {
			return errors.New("passport data corrupted")
		}
		// Get Field.
		field := passp.FieldByString(passFields[0])
		if field == passp.Uknown {
			return errors.New("passport data corrupted, field unknown")
		}
		passport.FillPassportData(field, passFields[1])
	}
	return nil
}
