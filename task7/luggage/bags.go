package luggage

import (
	"errors"
	"fmt"
	"sync"
)

type autoInc struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}

var ai autoInc // global instance

type Bag struct {
	serialNumber   int
	name           string
	additionalBags map[*Bag]int
	visited        bool
}

func CreateBag() *Bag {
	return &Bag{
		serialNumber:   ai.ID(),
		additionalBags: make(map[*Bag]int),
	}
}

// SetSerialNumber Number refers to the order the bag was discovered.
// 0 default - bag that is not discovered, yet.
func (b *Bag) SetSerialNumber(number int) {
	b.serialNumber = number
}

func (b *Bag) SerialNumber() int {
	return b.serialNumber
}

func (b *Bag) AddAdditionalBag(bag *Bag, amount int) {
	b.additionalBags[bag] = amount
}

func (b *Bag) AdditionalBags() *map[*Bag]int {
	return &b.additionalBags
}

func (b *Bag) PrintAdditionalBags() {
	for b := range b.additionalBags {
		fmt.Println(b)
	}
}

func (b *Bag) SetName(name string) {
	b.name = name
}

func (b *Bag) Name() string {
	return b.name
}

func (b *Bag) SetVisited(visited bool) {
	b.visited = visited
}

func (b *Bag) Visited() bool {
	return b.visited
}

// +++++++++++++++++++++++++BAGS++++++++++++++++++++++++++

// Exists returns *Bag if it is present in the slice of bags or nil if not.
func Exists(name string, bags *[]*Bag) (result *Bag) {
	result = nil
	for _, bag := range *bags {
		if bag.Name() == name {
			result = bag
			break
		}
	}
	return result
}

// AddUniqueBags will add bag that is unique to a given slice of bags.
func AddUniqueBags(bags *[]*Bag, bag *Bag) bool {
	if exists := Exists(bag.Name(), bags); exists == nil {
		*bags = append(*bags, bag)
		return true
	}
	return false
}

func BagBySerialNumber(id int, bags *[]*Bag) (result *Bag) {
	result = nil
	for _, bag := range *bags {
		if bag.SerialNumber() == id {
			result = bag
			break
		}
	}
	return result
}

func BagByName(name string, bags *[]*Bag) (*Bag, error) {
	for _, bag := range *bags {
		if bag.Name() == name {
			return bag, nil
		}
	}
	return CreateBag(), errors.New("no bag with that name found")
}

func PrintBags(bags *[]*Bag, additionalBags bool) {
	for _, bag := range *bags {
		fmt.Println("Main Bag", bag.SerialNumber(), bag.Name(), bag.Visited())
		if additionalBags {
			for ab := range *bag.AdditionalBags() {
				fmt.Println("-----Additional bags:", ab.SerialNumber(), ab.Name(), ab.Visited())
			}
		}
	}
}
