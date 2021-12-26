package datafile

import (
	"bufio"

	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/constants"
	"github.com/DagmarC/advent-of-code-2020/task7/luggage"
)

func LoadFileTask7() (*[]*luggage.Bag, error) {

	file, err := os.Open(constants.Task7)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	allBags := make([]*luggage.Bag, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		err := parseLineOfBags(line, &allBags)
		if err != nil {
			return nil, err
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return &allBags, nil
}

func parseLineOfBags(line string, allBags *[]*luggage.Bag) error {
	//line: shiny lime bags contain 3 muted magenta bags, 3 clear cyan bags.
	splittedText := strings.Split(strings.Trim(line, "\n"), " bags contain ")

	// 1st element in slice is the main bag - [shiny lime, 3 muted magenta bags, 3 clear cyan bags.]
	mainBag := manageMainBag(splittedText[0], allBags)

	// Bags the main bag can contain are additional bags. The splittedText[1]: 3 muted magenta bags, 3 clear cyan bags.
	err := parseAndAppendAdditionalBags(mainBag, splittedText[1], allBags)
	if err != nil {
		return err
	}
	return nil
}

// manageMainBag Will either create bag and add it to allBags slice of bags or just return an existing one.
func manageMainBag(name string, allBags *[]*luggage.Bag) *luggage.Bag {
	mainBag := luggage.Exists(name, allBags)
	if mainBag == nil {
		mainBag = luggage.CreateBag()
		mainBag.SetName(name)
		luggage.AddUniqueBags(allBags, mainBag)
	}
	//fmt.Println("MAIN", mainBag)
	return mainBag
}

func parseAndAppendAdditionalBags(mainBag *luggage.Bag, line string, allBags *[]*luggage.Bag) error {

	// line: 3 muted magenta bags, 3 clear cyan bags.
	reg := regexp.MustCompile(`([0-9]+) ([a-z]+ [a-z]+) [bags|bag]+[.|, ]*`)

	// loop over groups -  matches: [[3 muted magenta bags, 3, muted magenta], ...]
	matches := reg.FindAllStringSubmatch(line, -1)
	for _, m := range matches {
		//[3 muted magenta bags, 3, muted magenta],
		amountStr := m[1]
		name := m[2]

		// Convert number
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return err
		}

		// Create new and append to all Bags or obtain already existing one.
		additionalBag := luggage.Exists(name, allBags)
		if additionalBag == nil {
			additionalBag = luggage.CreateBag()
			additionalBag.SetName(name)
			luggage.AddUniqueBags(allBags, additionalBag)
		}
		//fmt.Println("ADD", additionalBag)
		mainBag.AddAdditionalBag(additionalBag, amount)

	}
	return nil
}
