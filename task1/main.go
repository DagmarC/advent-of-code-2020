package main

import (
	"fmt"
	"github.com/DagmarC/codeOfAdvent/task1/binarytree"
	"github.com/DagmarC/codeOfAdvent/task1/datafile"
	"github.com/DagmarC/codeOfAdvent/task1/solution"
)

func main() {
	fmt.Println("+++++++++++++++++TASK 1.1++++++++++++++++++++")

	// Task 1.1
	// Find the two entries that sum to 2020; what do you get if you multiply them together?
	tree := &binarytree.Tree{Root: nil}
	datafile.LoadFileIntoBinaryTree(tree)
	const sum int = 2020

	multipliedResult := tree.FindMultiplier(tree.Root, sum)
	if multipliedResult != 1 {
		fmt.Printf("Multiplication of x*y, where x+y=%d is %d\n", sum, multipliedResult)
	} else {
		fmt.Printf("There werent such two values x+y that were equal to %d.\n", sum)
	}
	fmt.Println("+++++++++++++++++TASK 1.2++++++++++++++++++++")
	// Task 1.2
	// In your expense report, what is the product of the three entries that sum to 2020?
	result, err := solution.Fin3dMultipliers(2020)
	if err != nil {
		fmt.Printf("There werent such three values x+y+z that were equal to %d.\n", sum)
	} else {
		fmt.Printf("Multiplication of x*y*z, where x+y+z=%d is %d\n", sum, result)
	}
}
