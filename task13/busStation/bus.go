package busStation

type Bus int

type DepartureTime int

type NearestBusDeparture struct {
	BusId               Bus
	NearestBusDeparture DepartureTime
	WaitTime            int
}

func (n *NearestBusDeparture) Update(id Bus, departure DepartureTime, waitTime int) {
	n.BusId = id
	n.NearestBusDeparture = departure
	n.WaitTime = waitTime
}

func (b *Bus) CalculateNextDepartureFrom(earliestDeparture DepartureTime) (DepartureTime, int) {
	reminder := int(earliestDeparture) % int(*b)

	if reminder == 0 {
		return earliestDeparture, 0
	}

	tmpDivision := earliestDeparture / DepartureTime(*b)
	nextBusDeparture := (tmpDivision + 1) * DepartureTime(*b)
	waitTime := nextBusDeparture - earliestDeparture
	return nextBusDeparture, int(waitTime)
}
