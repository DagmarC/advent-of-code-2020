package translation

import (
	"testing"

	"github.com/DagmarC/codeOfAdvent/utils"
)

var testTicketRules1 = TicketRules{
	{
		name:            "class",
		lowerBoundaries: [2]int{1, 5},
		upperBoundaries: [2]int{3, 7},
		availableFields: utils.MakeSet(),
	}, {
		name:            "row",
		lowerBoundaries: [2]int{6, 33},
		upperBoundaries: [2]int{11, 44},
		availableFields: utils.MakeSet(),
	}, {
		name:            "seat",
		lowerBoundaries: [2]int{13, 45},
		upperBoundaries: [2]int{40, 50},
		availableFields: utils.MakeSet(),
	},
}

var yourTicket = Ticket{
	[]int{11, 12, 13},
	make(map[string]int, 0),
}

var testTicketRules2 = TicketRules{
	{
		name:            "class",
		lowerBoundaries: [2]int{0, 4},
		upperBoundaries: [2]int{1, 19},
		availableFields: utils.MakeSet(),
	}, {
		name:            "row",
		lowerBoundaries: [2]int{0, 8},
		upperBoundaries: [2]int{5, 19},
		availableFields: utils.MakeSet(),
	}, {
		name:            "seat",
		lowerBoundaries: [2]int{0, 16},
		upperBoundaries: [2]int{13, 19},
		availableFields: utils.MakeSet(),
	},
}

var testNearbyTickets = []Ticket{
	{[]int{7, 3, 47},
		make(map[string]int, 0),
	},
	{[]int{40, 4, 50},
		make(map[string]int, 0),
	},
	{[]int{55, 2, 20},
		make(map[string]int, 0),
	},
	{[]int{38, 6, 12},
		make(map[string]int, 0),
	},
}

var testNearbyTickets2 = []Ticket{
	{[]int{3, 9, 18},
		make(map[string]int),
	},
	{[]int{15, 1, 5},
		make(map[string]int),
	},
	{[]int{5, 14, 9},
		make(map[string]int),
	},
}

func TestErrorRate(t *testing.T) {
	const expectedErrorRate int = 71
	actualErrorRate, _ := ErrorRate(&testNearbyTickets, &testTicketRules1)

	if actualErrorRate != expectedErrorRate {
		t.Errorf("Expected error rate is %d and actual is %d", expectedErrorRate, actualErrorRate)
		t.Fail()
	}
}

func TestTicketValidity(t *testing.T) {
	expectedErrorValues := []int{0, 4, 55, 12} // Error values

	for i, ticket := range testNearbyTickets {
		errorValue := ticket.Validate(&testTicketRules1)
		if errorValue != expectedErrorValues[i] {
			t.Errorf("Expected error value is %d and actual is %d", expectedErrorValues[i], errorValue)
			t.Fail()
		}
	}
}

func TestTicketOrdering(t *testing.T) {
	// Prepare mocks.
	expectedRulesOrdering := map[string]int{
		"row":   0,
		"class": 1,
		"seat":  2,
	}

	TicketRulesOrdering(&testTicketRules2, &testNearbyTickets2, &yourTicket)

	for k, v := range yourTicket.Rules {
		if expectedRulesOrdering[k] != v {
			t.Fail()
			t.Logf("Actual %v and expected %v\n", v, expectedRulesOrdering[k])

		}
	}
	t.Logf("Actual %v and expected %v\n", yourTicket.Rules, expectedRulesOrdering)
}
