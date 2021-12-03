package graphsearch

import (
	"github.com/DagmarC/codeOfAdvent/task7/luggage"
	"log"
)

// CreateOppositeBags If X -> A, B (X bag has additional bags A and B), then the resultant slice will contain
// A -> X, B -> X. It means that all arrows will be inverted from IN to OUT.
// Need to create a copy of all bags to not interfere the original slice. All new bags will have the same name and the
// IDs will be the same after subtraction of the length of the slice -> NewID = original ID + length(slice)
func CreateOppositeBags(allBags *[]*luggage.Bag) []*luggage.Bag {

	oppositeBags := make([]*luggage.Bag, 0) // new struct
	// Create the copy of the bags and slice -> name is the same but IDs are greater by length of the original slice.
	for _, bag := range *allBags {
		newBag := luggage.CreateBag()
		newBag.SetName(bag.Name())
		oppositeBags = append(oppositeBags, newBag)
	}

	// Loop over the original allBags and additionalBags and then obtain the bags from oppositeBags via name.
	// To make it opposite, loop over all bags and for each additional bag attach the mainBag.
	// Original: Main bag -> Original bag
	// Opposite: Additional Bag -> Main Bag
	for _, bag := range *allBags {
		// Get tha bag with the same name but from opposite directions.
		bagOpposite, err := luggage.BagByName(bag.Name(), &oppositeBags)
		if err != nil {
			log.Fatal(err)
		}
		for ab := range *bag.AdditionalBags() {
			additionalOppositeBag, err := luggage.BagByName(ab.Name(), &oppositeBags)
			if err != nil {
				log.Fatal(err)
			}
			// Assign the bagOpposite to additional bags from oppositeBags.
			additionalOppositeBag.AddAdditionalBag(bagOpposite, 1)
		}
	}
	return oppositeBags
}

// DepthSearch TASK 1: How many bag colors can eventually contain at least one shiny gold bag?
func DepthSearch(initialIndexBag *luggage.Bag) int {
	return depthSearchRec(initialIndexBag, 0)
}

// depthSearchRec TASK 1
func depthSearchRec(indexBag *luggage.Bag, count int) int {
	indexBag.SetVisited(true)
	//fmt.Println("Recursion", indexBag)

	for vertexBag := range *indexBag.AdditionalBags() {
		// If not visited -> continue with searching on this vertex.
		if !vertexBag.Visited() {
			count = depthSearchRec(vertexBag, count+1)
		}
	}
	return count
}

// DepthSearchWeightedEdges TASK 2: How many individual bags are required inside your single shiny gold bag?
func DepthSearchWeightedEdges(initialIndexBag *luggage.Bag) int {
	return depthSearchWeightedEdgesRec(initialIndexBag)
}

func depthSearchWeightedEdgesRec(indexBag *luggage.Bag) int {
	edgeSum := 0
	for vertexBag, edgeValue := range *indexBag.AdditionalBags() {
		edgeSum += edgeValue + edgeValue*depthSearchWeightedEdgesRec(vertexBag)
	}
	return edgeSum
}
