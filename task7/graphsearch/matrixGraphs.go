package graphsearch

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/task7/luggage"
)

// CreateOppositeBags If X -> A, B (X bag has additional bags A and B), then the resultant slice will contain
// A -> X, B -> X. It means that all arrows will be inverted from IN to OUT.
// Need to create a copy of all bags to not interfere the original slice. All new bags will have the same name and the
// IDs will be the same after subtraction of the length of the slice -> NewID = original ID + length(slice)
func CreateOppositeBags(allBags *[]*luggage.Bag) []*luggage.Bag {

	oppositeBags := make([]*luggage.Bag, 0) // new struct
	// Create the copy of the bags and slice -> Name is the same but IDs are greater by length of the original slice.
	for _, bag := range *allBags {
		newBag := luggage.CreateBag()
		newBag.SetName(bag.GetName())
		oppositeBags = append(oppositeBags, newBag)
	}

	// Loop over the original allBags and additionalBags and then obtain the bags from oppositeBags via name.
	// To make it opposite, loop over all bags and for each additional bag attach the mainBag.
	// Original: Main bag -> Original bag
	// Opposite: Additional Bag -> Main Bag
	for _, bag := range *allBags {
		// Get tha bag with the same name but from opposite directions.
		bagOpposite := luggage.GetByName(bag.GetName(), &oppositeBags)
		for ab := range *bag.GetAdditionalBags() {
			additionalOppositeBag := luggage.GetByName(ab.GetName(), &oppositeBags)
			// Assign the bagOpposite to additional bags from oppositeBags.
			additionalOppositeBag.AddAdditionalBag(bagOpposite, 1)
		}
	}
	return oppositeBags
}

func DepthSearch(initialIndexBag *luggage.Bag) int {
	return DepthSearchRec(initialIndexBag, 0)
}

func DepthSearchRec(indexBag *luggage.Bag, count int) int {
	indexBag.SetVisited(true)
	fmt.Println("Recursion", indexBag)

	for vertexBag := range *indexBag.GetAdditionalBags() {
		// If not visited -> continue with searching on this vertec
		if !vertexBag.GetVisited() {
			count = DepthSearchRec(vertexBag, count+1)
		}
	}
	return count
}
