package main

import (
	"fmt"
	"sort"

	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task16/translation"
)

func main() {
	rules, yourTicket, nearbyTickets, departureKeys := datafile.LoadDay16()

	// TASK 1
	errorRate, wrongTicketsIDs := translation.ErrorRate(nearbyTickets, rules)
	fmt.Println("What is your ticket scanning error rate?", errorRate)

	// TASK 2
	deleteInvalidTickets(nearbyTickets, wrongTicketsIDs)

	translation.TicketRulesOrdering(rules, nearbyTickets, yourTicket)
	fmt.Println(yourTicket)

	fmt.Println("Word starts with departure. What do you get if you multiply those six values together?",
		departureValuesMultiplied(yourTicket, departureKeys))
}

// departureValuesMultiplied goes through all departure keys, obtain the position of the key in the ticket via
// ticker.Rules[ds] and the gets the value at given position on your ticket within ticket IDs. Multiplies it altogether.
func departureValuesMultiplied(ticket *translation.Ticket, departureKeys *[]string) int {
	result := 1
	for _, dk := range *departureKeys {
		result *= ticket.IDs[ticket.Rules[dk]]
	}
	return result
}

func deleteInvalidTickets(nearbyTickets *[]translation.Ticket, invalidIDs []int) {
	sort.Ints(invalidIDs)

	for i := len(invalidIDs) - 1; i >= 0; i-- {
		*nearbyTickets = append((*nearbyTickets)[:invalidIDs[i]], (*nearbyTickets)[invalidIDs[i]+1:]...)
	}
}
