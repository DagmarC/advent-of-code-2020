// In the above rules, the following options would be available to you:
//
//A bright white bag, which can hold your shiny gold bag directly.
//A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
//A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
//A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
//So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.
package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task7/graphsearch"
	"github.com/DagmarC/codeOfAdvent/task7/luggage"
	"log"
)

func main() {
	//lineTest := "shiny lime bags contain 3 muted magenta bags, 3 clear cyan bags, 7 muted magenta bags.\n"
	//allBagsTest := make([]*luggage.Bag, 0)
	allBags, err := datafile.LoadFileTask7()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("__________________SOLUTION TASK 1______________")
	oppositeBags := graphsearch.CreateOppositeBags(allBags)
	fmt.Println()
	shinyGoldBagOpposite, err := luggage.BagByName("shiny gold", &oppositeBags)
	if err != nil {
		log.Fatal(err)
	}
	count := graphsearch.DepthSearch(shinyGoldBagOpposite)
	fmt.Println("RESULT TASK 1: ", count)

	shinyGoldBag, err := luggage.BagByName("shiny gold", allBags)
	if err != nil {
		log.Fatal(err)
	}
	result := graphsearch.DepthSearchWeightedEdges(shinyGoldBag)
	fmt.Println("RESULT TASK 2: ", result)

}
