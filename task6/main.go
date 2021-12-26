package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/DagmarC/advent-of-code-2020/datafile"
)

func main() {
	allUniqueAnswers, allAnswers, err := datafile.LoadFileTask6()
	if err != nil {
		log.Fatal(err)
	}
	// TASK 1
	fmt.Println("Count of All \"yes\" unique answers anyone answers in each group:", countAnswersAnyone(allUniqueAnswers))
	// TASK 2
	fmt.Println("Count of All \"yes\" answers thath everyone answers "+
		"in each group:", countAnswersEveryone(allAnswers))
}

// TASK 1 - Anyone answers yes
func countAnswersAnyone(allAnswers *map[int]string) int {
	sum := 0
	for _, answers := range *allAnswers {
		sum += len(answers)
	}
	return sum
}

// TASK 2 - All must have answered yes.
func countAnswersEveryone(allAnswers *map[int][]string) int {
	sumAnswers := 0
	// answers : [a, ab, axb]
	for _, answers := range *allAnswers {
		// SORT answers per person by length.
		sort.Slice(answers, func(i, j int) bool {
			return len(answers[i]) < len(answers[j])
		})
		// INTERSECT answers in each group
		sumAnswers += findAnswersIntersection(&answers)
	}
	return sumAnswers
}

// findAnswersIntersection Slice of answers [a-z]. Each person`s answers is represented via one element in slice.
// Function result example:
// input: sorted by length [b, bc, abz] -> 3 separate people
// output: 1 -> b is the only answers common for everyone.
func findAnswersIntersection(answers *[]string) int {
	if len(*answers) == 0 {
		return 0
	}
	// Everyone (all 1 person) answered "yes"
	if len(*answers) == 1 {
		return len((*answers)[0])
	}

	firstPersonAnswers := (*answers)[0]
	sum := 0
	// You have to check all people`s answers for each letter to count it as correct answer.
	// Example: answer: letter a
answersCheck:
	for _, answer := range firstPersonAnswers {
		for i := 1; i < len(*answers); i++ {
			// Example: if abc.contains(a)
			if !strings.Contains((*answers)[i], string(answer)) {
				continue answersCheck
			}
		}
		// At this point we can assume that all other people answers 'yes' to current answer.
		sum++
	}
	return sum
}
