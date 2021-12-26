// https://adventofcode.com/2020/day/10#part2
// Each of your joltage adapters is rated for a specific output joltage (your puzzle input). Any given adapter can take
//an input 1, 2, or 3 jolts lower than its rating and still produce its rated output joltage.
//In addition, your device has a built-in joltage adapter rated for 3 jolts higher than the highest-rated adapter in your bag.
package main

import (
	"fmt"
	"log"

	"github.com/DagmarC/advent-of-code-2020/datafile"
	"github.com/DagmarC/advent-of-code-2020/task10/adapterpkg"
)

func main() {
	allAdapters, err := datafile.LoadFileTask10()
	if err != nil {
		log.Fatal(err)
	}
	joltDifferences, err := adapterpkg.DistributeAdapters(&allAdapters)
	if err != nil {
		log.Fatal(err)
	}
	// What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?
	sizeOneDistribution := len(joltDifferences[adapterpkg.One])
	sizeThreeDistribution := len(joltDifferences[adapterpkg.Three])

	fmt.Println("RESULT of TASK 1: ", sizeOneDistribution*sizeThreeDistribution)

	// What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?
	allCombinations := adapterpkg.GetAllAdaptersCombinations(&allAdapters)
	fmt.Println("Total number of distinct ways you can arrange the adapters:", allCombinations)
}
