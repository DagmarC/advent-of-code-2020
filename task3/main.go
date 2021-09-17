// Determine the number of trees you would encounter if, for each of the following slopes, you start
// at the top-left corner and traverse the map all the way to the bottom:
//
//Right 1, down 1.
//Right 3, down 1. (This is the slope you already checked.)
//Right 5, down 1.
//Right 7, down 1.
//Right 1, down 2.
//In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively;
//multiplied together, these produce the answer 336.
package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/datafile"
)

const treeField string = "#"

//const treeVisited string = "X"
//const spaceVisited string = "O"

func main() {
	forest := datafile.LoadAndParseTask3()

	countTree := countTrees(&forest, 1, 1)
	fmt.Println("Number of trees visited down the road - first traversal: ", countTree)

	countTree *= countTrees(&forest, 3, 1)
	//fmt.Println("Number of trees visited down the road: ", countTree)

	countTree *= countTrees(&forest, 5, 1)
	//fmt.Println("Number of trees visited down the road: ", countTree)

	countTree *= countTrees(&forest, 7, 1)
	//fmt.Println("Number of trees visited down the road: ", countTree)

	countTree *= countTrees(&forest, 1, 2)
	fmt.Println("Number of trees visited down the road and multiplied together after all traversals: ", countTree)

	//fmt.Println("*********PRINT***********")
	//for i, _ := range forest {
	//	for _, row := range forest[i]{
	//	fmt.Print(row)
	//}
	//	fmt.Println()
	//}
}

func countTrees(forest *[][]string, right, down int) int {

	countTree := 0
	rowIndex := 0

	lenLine := len((*forest)[0])

	for i := 0; i < len(*forest); i += down {

		if rowIndex >= lenLine {
			rowIndex %= lenLine
		}
		if (*forest)[i][rowIndex] == treeField {
			countTree++
			//(*forest)[i][rowIndex] = treeVisited
		} else {
			//(*forest)[i][rowIndex] = spaceVisited
		}
		rowIndex += right
	}
	return countTree
}
