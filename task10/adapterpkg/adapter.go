package adapterpkg

import "errors"

type Jolt int

type Adapter Jolt

const (
	One Jolt = iota + 1
	Two
	Three
	Undefined
)

// JoltAdaptersMap adapters can only connect to a source 1-3 jolts lower than its rating.
// E.g. map[One][1, 2, ... ] --> means that to numbers 1 and 2 there exists adapters 1+1 2+1 in the input array
// to connect the charging outlet.
type JoltAdaptersMap map[Jolt][]Adapter

func (a Adapter) Difference(jolt Adapter) Jolt {
	switch jolt - a {
	case 1:
		return One
	case 2:
		return Two
	case 3:
		return Three
	default:
		return Undefined
	}
}

func (a *JoltAdaptersMap) InitializeDistribution() {
	(*a)[One] = make([]Adapter, 0)
	(*a)[Two] = make([]Adapter, 0)
	(*a)[Three] = make([]Adapter, 0)
}

func (a *JoltAdaptersMap) InitializeAdaptersMap(adapters *[]Adapter) error {
	if len(*adapters) == 0 {
		return errors.New("no adapters given")
	}

	for _, adapter := range *adapters {
		(*a)[Jolt(adapter)] = make([]Adapter, 0)
	}
	return nil
}

type visitedSign int

const (
	grey  visitedSign = 0
	white visitedSign = -1
)

func visitedAdapters(adapters *[]Adapter) (map[Adapter]visitedSign, error) {
	visited := make(map[Adapter]visitedSign, 0)

	if len(*adapters) == 0 {
		return nil, errors.New("no adapters given")
	}

	for _, adapter := range *adapters {
		visited[adapter] = white
	}
	return visited, nil
}
