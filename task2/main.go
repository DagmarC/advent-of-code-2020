// Task: Each line gives the password policy and then the password. The password policy indicates the lowest
//and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means
//that the password must contain a at least 1 time and at most 3 times.
package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"log"
)

func main() {
	passwordDefinitions, err := datafile.LoadFileTask2()
	if err != nil {
		log.Fatal(err)
	}
	countValidPasswords := 0
	for _, pd := range passwordDefinitions {
		if pd.ValidatePassword() {
			countValidPasswords++
		}
	}
	fmt.Println("Number of valid passwords: ", countValidPasswords)
}
