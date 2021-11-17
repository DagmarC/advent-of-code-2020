package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/codeOfAdvent/datafile"
	"github.com/DagmarC/codeOfAdvent/task14/initprogram"
)

// https://adventofcode.com/2020/day/14
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
