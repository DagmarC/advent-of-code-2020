package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/task14/initprogram"
)

// https://adventofcode.com/2020/day/14
// https://github.com/colinodell/advent-2020/blob/3b889146ced9248713826915d0efff2b86a37135/day14/day14.go#L98 Good solution
func main() {
	//test()
	instructions, err := datafile.LoadFileAndRunTask14(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("T", instructions.MemLocations())
	fmt.Println(" What is the sum of all values left in memory after it completes? Task 1", instructions.SumOfMemLocations())

	instructions, err = datafile.LoadFileAndRunTask14(2)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("T", instructions.MemLocations())
	fmt.Println(" What is the sum of all values left in memory after it completes? Task 2", instructions.SumOfMemLocations())
}

func test() {
	instructions := initprogram.InitializeInitProgram()
	instructions.UpdateBitmask("0X1X110000X00X001000011000100110X011")
	instructions.Write(1, 745799437)
	fmt.Println(instructions.MemLocations())
}
