package main

import (
	"fmt"

	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/task17/cubes"
	"github.com/DagmarC/advent-of-code-2020/utils"
)

func main() {
	file := datafile.ReadLines(utils.GetInputPath())
	c := &cubes.Cubes{}
	c.Initialize(file)

	cycles := 6
	for cycles > 0 {
		fmt.Println("BOOT PROCESS n:", cycles)
		c = cubes.BootProcess(*c)
		cycles--
	}
	fmt.Println("------TASK1------")
	fmt.Println("How many cubes are left in the active state after the sixth cycle?", c.Active())

}
