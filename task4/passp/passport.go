package passp

import (
	"regexp"
	"strconv"
)

type Field int64

const (
	Byr Field = iota
	Iyr
	Eyr
	Hgt
	Hcl
	Ecl
	Pid
	Cid
	Uknown
)

type Passport struct {
	data            map[string]string
	InitFieldsCount int
}

func CreatePassport() *Passport {
	return &Passport{data: make(map[string]string)}
}

func (pp *Passport) FillPassportData(passportField Field, data string) {
	pp.data[passportField.String()] = data
	pp.InitFieldsCount++
}

func (pp *Passport) GetPassportsData() map[string]string {
	return pp.data
}

func (f Field) String() string {
	switch f {
	case Byr:
		return "byr"
	case Iyr:
		return "iyr"
	case Eyr:
		return "eyr"
	case Hgt:
		return "hgt"
	case Hcl:
		return "hcl"
	case Ecl:
		return "ecl"
	case Pid:
		return "pid"
	case Cid:
		return "cid"
	}
	return "unknown"
}

func GetField(field string) Field {
	switch field {
	case "byr":
		return Byr
	case "iyr":
		return Iyr
	case "eyr":
		return Eyr
	case "hgt":
		return Hgt
	case "hcl":
		return Hcl
	case "ecl":
		return Ecl
	case "pid":
		return Pid
	case "cid":
		return Cid
	}
	return Uknown
}

// Valid Passports - Part II.
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.
//cid (Country ID) - ignored, missing or not.

func (pp *Passport) ValidatePassport() bool {
	for key, data := range pp.data {
		isValid := pp.validateData(data, GetField(key))
		// All data must be valid at the same time.
		if !isValid {
			return false
		}
	}
	return true
}

func (pp *Passport) validateData(data string, field Field) bool {
	switch field {
	case Byr:
		return validateYearInput(data, 4, 1920, 2002)
	case Iyr:
		return validateYearInput(data, 4, 2010, 2020)
	case Eyr:
		return validateYearInput(data, 4, 2020, 2030)
	case Hgt:
		return validateHeightInput(data)
	case Hcl:
		return validateHairColor(data)
	case Ecl:
		return validateEyeColor(data)
	case Pid:
		return validatePassportID(data)
	case Cid:
		return true // Optional field.
	}
	return false
}

//pid (Passport ID) - a nine-digit number, including leading zeroes.
func validatePassportID(data string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", data)
	return match
}

//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validateEyeColor(data string) bool {
	match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", data)
	return match
}

//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validateHairColor(data string) bool {
	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", data)
	return match
}

//hgt (Height) - a number followed by either cm or in:
func validateHeightInput(data string) bool {
	match, _ := regexp.MatchString("^[0-9]+(cm|in)+$", data)
	if !match {
		return false
	}

	heightRegex := regexp.MustCompile("^[0-9]+")
	heightStr := heightRegex.FindAllString(data, -1)
	if len(heightStr) != 1 {
		return false
	}

	height, err := strconv.Atoi(heightStr[0])
	if err != nil {
		return false
	}

	metricsRegex := regexp.MustCompile("(cm|in)$")
	metricsStr := metricsRegex.FindAllString(data, -1)
	if len(metricsStr) != 1 {
		return false
	}
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	if metricsStr[0] == "cm" {
		return height >= 150 && height <= 193
	} else {
		return height >= 59 && height <= 76
	}
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validateYearInput(data string, digits int, lowerBoundary, upperBoundary int) bool {
	if len(data) != digits {
		return false
	}
	year, err := strconv.Atoi(data)
	if err != nil {
		return false
	}
	return year >= lowerBoundary && year <= upperBoundary
}
