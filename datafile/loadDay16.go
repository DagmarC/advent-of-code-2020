package datafile

import (
	"strings"

	"github.com/DagmarC/codeOfAdvent/constants"
	"github.com/DagmarC/codeOfAdvent/task16/translation"
)

const rulesStr string = "rules:"
const yourTicketStr string = "your ticket:"
const nearbyTicketsStr string = "nearby tickets:"

func LoadDay16() (*translation.TicketRules, *translation.Ticket, *[]translation.Ticket, *[]string) {

	var rules translation.TicketRules
	var nearbyTickets []translation.Ticket
	var yourTicket *translation.Ticket
	var departureKeys = make([]string, 0)

	lines := ReadLines(constants.Task16)

	loader := rulesStr // The default one.
	for _, line := range lines {
		// decide what to load
		if loader != nearbyTicketsStr { // if loader is nearbyTicketsStr, no other switch is expected.
			if ok := switchLoader(&loader, line); ok {
				continue // skip this line
			}
		}

		switch loader {
		case rulesStr:
			rules.Add(line, &departureKeys)
		case yourTicketStr:
			yourTicket = translation.CreateTicket(line)
		case nearbyTicketsStr:
			nearbyTickets = append(nearbyTickets, *translation.CreateTicket(line))
		}
	}
	return &rules, yourTicket, &nearbyTickets, &departureKeys
}

// switch loader changes the loader if needed. If it returns true, this line will be skipped in loading.
func switchLoader(loader *string, line string) bool {
	skipLine := false

	if strings.Contains(line, yourTicketStr) {
		*loader = yourTicketStr
		skipLine = true

	} else if strings.Contains(line, nearbyTicketsStr) {
		*loader = nearbyTicketsStr
		skipLine = true

	} else if line == "" {
		skipLine = true
	}

	return skipLine
}
