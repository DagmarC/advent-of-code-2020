package translation

import (
	"errors"
	"regexp"
	"sort"
	"strings"

	"github.com/DagmarC/codeOfAdvent/utils"
)

type Ticket struct {
	IDs   []int
	Rules map[string]int
}

func CreateTicket(line string) *Ticket {
	numbers := make([]int, 0)

	numbersStr := strings.Split(line, ",")
	for _, n := range numbersStr {
		numbers = append(numbers, utils.ToInt(n))
	}
	return &Ticket{
		IDs:   numbers,
		Rules: make(map[string]int, 0),
	}
}

func (t *Ticket) Validate(rules *TicketRules) (errorValue int) {

	for _, number := range t.IDs {
		if ok := rules.inRange(number); !ok {
			errorValue += number
		}
	}
	return errorValue
}

// -------------------------------------------------

type TicketRule struct {
	name            string
	lowerBoundaries [2]int
	upperBoundaries [2]int
	availableFields *utils.CustomSet
}

func (tr *TicketRule) Name() string {
	return tr.name
}

func (tr *TicketRule) Fields() *utils.CustomSet {
	return tr.availableFields
}

func (tr *TicketRule) inRange(n int) bool {
	for i, lb := range tr.lowerBoundaries {
		if lb <= n && n <= tr.upperBoundaries[i] {
			return true
		}
	}
	return false
}

type TicketRules []TicketRule

func CreateTicketRule(line string) *TicketRule {

	reg := regexp.MustCompile("([a-z]+.[a-z]*): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
	match := reg.FindStringSubmatch(line)

	return &TicketRule{
		match[1],
		[2]int{utils.ToInt(match[2]), utils.ToInt(match[4])},
		[2]int{utils.ToInt(match[3]), utils.ToInt(match[5])},
		utils.MakeSet(),
	}
}

func (r *TicketRules) Add(line string, departureKeys *[]string) {
	ticketRule := CreateTicketRule(line)

	if strings.Contains(ticketRule.Name(), "departure") {
		*departureKeys = append(*departureKeys, ticketRule.Name())
	}
	*r = append(*r, *ticketRule)
}

// inRange checks if the number n is in range for any of the rules.
func (r *TicketRules) inRange(n int) bool {
	for _, rule := range *r {
		if ok := rule.inRange(n); ok {
			return ok
		}
	}
	return false
}

func (r *TicketRules) GetByName(name string) (TicketRule, error) {
	for _, rule := range *r {
		if rule.Name() == name {
			return rule, nil
		}
	}
	return TicketRule{}, errors.New("no rule found")
}

// -------------------------------------------------

type ByFields TicketRules

func (s ByFields) Len() int {
	return len(s)
}

func (s ByFields) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByFields) Less(i, j int) bool {
	return s[i].availableFields.Size() < s[j].availableFields.Size()
}

// -------------------------------------------------

// ErrorRate identifies and sums up all numbers in tickets that are not in range from any of the rules.
func ErrorRate(tickets *[]Ticket, rules *TicketRules) (int, []int) {
	errorRateAll := 0
	wrongTickets := make([]int, 0)

	for i, ticket := range *tickets {
		errorV := ticket.Validate(rules)
		if errorV != 0 {
			errorRateAll += errorV
			wrongTickets = append(wrongTickets, i)
		}
	}
	return errorRateAll, wrongTickets
}

// ----------------------Task2-----------------------

func TicketRulesOrdering(rules *TicketRules, nearbyTickets *[]Ticket, yourTicket *Ticket) {
	ticketIDsLen := len((*nearbyTickets)[0].IDs)

	for _, rule := range *rules {
		for pos := 0; pos < ticketIDsLen; pos++ {
			assignRulePositions(pos, &rule, nearbyTickets)
		}
	}

	getFieldRec(rules, yourTicket)
	return
}

// assignRulePositions Valid position for a rule is when All tickets at given position are in range for the given rule.
func assignRulePositions(pos int, rule *TicketRule, tickets *[]Ticket) {
	for _, t := range *tickets {
		if ok := rule.inRange(t.IDs[pos]); !ok {
			return // Do not assign if at least one ticket is not valid.
		}
	}
	rule.availableFields.Add(pos) // Valid position
}

func getFieldRec(rules *TicketRules, ticket *Ticket) {

	if len(*rules) == 0 {
		return
	}

	// 1. Sort to get the rule with only 1 available field.
	sort.Sort(ByFields(*rules))

	// 2. Get the 1st available field.
	firstRule := (*rules)[0]
	field := firstRule.availableFields.First()

	// 3. Set your ticket - assign field to rule name.
	ticket.Rules[firstRule.Name()] = field

	// 4. Delete the first rule from rules - already assigned.
	*rules = (*rules)[1:]

	// 5. Delete the field from all rules.
	deletePositionFromRules(rules, field)

	// 5. Repeat.
	getFieldRec(rules, ticket)
}

func deletePositionFromRules(rules *TicketRules, position int) {
	for _, rule := range *rules {
		if ok := rule.availableFields.Exists(position); ok {
			err := rule.availableFields.Remove(position)
			utils.Check(err)
		}
	}
}
