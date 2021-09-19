// Count the number of valid passports - those that have all required fields - all except CID.
// Treat cid as optional. In your batch file, how many passports are valid?
package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task4/passp"
)

func main() {
	passports := datafile.LoadAndParseTask4()
	fmt.Println("-------MAIN--------")
	// Treat cid as optional.
	validPassportsCount := countValidPasswords(&passports)
	fmt.Println("Number of valid passwords:", validPassportsCount)
}

// countValidPasswords Part I Task: will count all valid passwords -> Either all 8 fields have been initialised or CID as optional
// was not initialised --> only 7 fields have been initialised.
func countValidPasswords(passports *[]*passp.Passport) int {
	count := 0
	for _, p := range *passports {
		_, cidInit := p.GetPassportsData()[passp.Cid.String()]

		if p.InitFieldsCount == 8 {
			if valid := p.ValidatePassport(); valid {
				count++
			}
		} else if p.InitFieldsCount == 7 && !cidInit {
			// if ok is false, the CID is not initialized, which is okay, since it is optional.
			if valid := p.ValidatePassport(); valid {
				fmt.Println(p)
				count++
			}
		}
	}
	return count
}
