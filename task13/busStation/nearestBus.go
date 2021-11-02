package busStation

type BusId int

type DepartureTime int

type NearestBusDeparture struct {
	BusId               BusId
	NearestBusDeparture DepartureTime
	Offset              int
}

func (n *NearestBusDeparture) Update(id BusId, departure DepartureTime, waitTime int) {
	n.BusId = id
	n.NearestBusDeparture = departure
	n.Offset = waitTime
}

// ByBusId CUSTOM SORT BY BUS ID -> IMPLEMENT SORT INTERFACE
type ByBusId []NearestBusDeparture

func (a ByBusId) Len() int {
	return len(a)
}

func (a ByBusId) Less(i, j int) bool {
	return a[i].BusId > a[j].BusId
}

func (a ByBusId) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (b *BusId) CalculateNextDepartureFrom(earliestDeparture DepartureTime) (DepartureTime, int) {
	reminder := int(earliestDeparture) % int(*b)

	if reminder == 0 {
		return earliestDeparture, 0
	}

	tmpDivision := earliestDeparture / DepartureTime(*b)
	nextBusDeparture := (tmpDivision + 1) * DepartureTime(*b)
	waitTime := nextBusDeparture - earliestDeparture
	return nextBusDeparture, int(waitTime)
}
