package adapterpkg

import (
	"fmt"
	"log"
)

func GetAllAdaptersCombinations(allAdapters *[]Adapter) int {

	adaptersGraph, err := createAdapterArrangements(allAdapters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Graph: ", adaptersGraph)

	visitedAdapters, err := visitedAdapters(allAdapters)
	if err != nil {
		log.Fatal(err)
	}

	getAllAdaptersCombinationsRec(&adaptersGraph, (*allAdapters)[0], &visitedAdapters)
	fmt.Println("Visited", visitedAdapters)

	return int(visitedAdapters[0])
}

func getAllAdaptersCombinationsRec(graph *JoltAdaptersMap, currentAdapter Adapter, visitedAdapters *map[Adapter]visitedSign) visitedSign {

	(*visitedAdapters)[currentAdapter] = grey
	adjacentVertices := (*graph)[Jolt(currentAdapter)]

	// grey == 0
	sumCurrentAdapter := grey
	for _, vertex := range adjacentVertices {
		// white == -1 --> unvisited
		if (*visitedAdapters)[vertex] == white {
			getAllAdaptersCombinationsRec(graph, vertex, visitedAdapters)
			sumCurrentAdapter += (*visitedAdapters)[vertex]

		} else {
			sumCurrentAdapter += (*visitedAdapters)[vertex]
		}
		(*visitedAdapters)[currentAdapter] = sumCurrentAdapter
	}

	if len(adjacentVertices) == 0 {
		(*visitedAdapters)[currentAdapter] = 1
	}

	return sumCurrentAdapter
}
